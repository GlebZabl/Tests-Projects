package middlewares

// this file is moderated functions from https://github.com/labstack/echo
import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type (
	BindUnmarshaler interface {
		UnmarshalParam(param string) error
	}
)

func Bind(i interface{}, req http.Request) (err error) {
	if req.ContentLength == 0 {
		if req.Method == http.MethodGet || req.Method == http.MethodDelete {
			if err = bindData(i, getQueryParams(req), "query"); err != nil {
				return errors.New("")
			}
			return
		}
		return errors.New("")
	}
	ctype := req.Header.Get("Content-Type")
	switch {
	case strings.HasPrefix(ctype, "application/json"):
		if err = json.NewDecoder(req.Body).Decode(i); err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				return errors.New("")
			} else if _, ok := err.(*json.SyntaxError); ok {
				return errors.New("")
			} else {
				return errors.New("")
			}
			return errors.New("")
		}
	case strings.HasPrefix(ctype, "application/json"), strings.HasPrefix(ctype, "text/xml"):
		if err = xml.NewDecoder(req.Body).Decode(i); err != nil {
			if _, ok := err.(*xml.UnsupportedTypeError); ok {
				return errors.New("")
			} else if _, ok := err.(*xml.SyntaxError); ok {
				return errors.New("")
			} else {
				return errors.New("")
			}
			return errors.New("")
		}
	case strings.HasPrefix(ctype, "application/json"), strings.HasPrefix(ctype, "multipart/form-data"):
		params, err := getFormParams(req)
		if err != nil {
			return errors.New("")
		}
		if err = bindData(i, params, "form"); err != nil {
			return errors.New("")
		}
	default:
		return errors.New("")
	}
	return
}

func bindData(ptr interface{}, data map[string][]string, tag string) error {
	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()

	if typ.Kind() != reflect.Struct {
		return errors.New("binding element must be a struct")
	}

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}
		structFieldKind := structField.Kind()
		inputFieldName := typeField.Tag.Get(tag)

		if inputFieldName == "" {
			inputFieldName = typeField.Name
			if _, ok := bindUnmarshaler(structField); !ok && structFieldKind == reflect.Struct {
				if err := bindData(structField.Addr().Interface(), data, tag); err != nil {
					return err
				}
				continue
			}
		}

		inputValue, exists := data[inputFieldName]
		if !exists {
			inputFieldName = strings.ToLower(inputFieldName)
			for k, v := range data {
				if strings.ToLower(k) == inputFieldName {
					inputValue = v
					exists = true
					break
				}
			}
		}

		if !exists {
			continue
		}

		if ok, err := unmarshalField(typeField.Type.Kind(), inputValue[0], structField); ok {
			if err != nil {
				return err
			}
			continue
		}

		numElems := len(inputValue)
		if structFieldKind == reflect.Slice && numElems > 0 {
			sliceOf := structField.Type().Elem().Kind()
			slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
			for j := 0; j < numElems; j++ {
				if err := setWithProperType(sliceOf, inputValue[j], slice.Index(j)); err != nil {
					return err
				}
			}
			val.Field(i).Set(slice)
		} else if err := setWithProperType(typeField.Type.Kind(), inputValue[0], structField); err != nil {
			return err

		}
	}
	return nil
}

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	if ok, err := unmarshalField(valueKind, val, structField); ok {
		return err
	}

	switch valueKind {
	case reflect.Ptr:
		return setWithProperType(structField.Elem().Kind(), val, structField.Elem())
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32, structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	default:
		return errors.New("unknown type")
	}
	return nil
}

func unmarshalField(valueKind reflect.Kind, val string, field reflect.Value) (bool, error) {
	switch valueKind {
	case reflect.Ptr:
		return unmarshalFieldPtr(val, field)
	default:
		return unmarshalFieldNonPtr(val, field)
	}
}

func bindUnmarshaler(field reflect.Value) (BindUnmarshaler, bool) {
	ptr := reflect.New(field.Type())
	if ptr.CanInterface() {
		iface := ptr.Interface()
		if unmarshaler, ok := iface.(BindUnmarshaler); ok {
			return unmarshaler, ok
		}
	}
	return nil, false
}

func unmarshalFieldNonPtr(value string, field reflect.Value) (bool, error) {
	if unmarshaler, ok := bindUnmarshaler(field); ok {
		err := unmarshaler.UnmarshalParam(value)
		field.Set(reflect.ValueOf(unmarshaler).Elem())
		return true, err
	}
	return false, nil
}

func unmarshalFieldPtr(value string, field reflect.Value) (bool, error) {
	if field.IsNil() {
		field.Set(reflect.New(field.Type().Elem()))
	}
	return unmarshalFieldNonPtr(value, field.Elem())
}

func setIntField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	intVal, err := strconv.ParseInt(value, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	uintVal, err := strconv.ParseUint(value, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(value string, field reflect.Value) error {
	if value == "" {
		value = "false"
	}
	boolVal, err := strconv.ParseBool(value)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0.0"
	}
	floatVal, err := strconv.ParseFloat(value, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func getQueryParams(r http.Request) url.Values {
	return r.URL.Query()
}

func getFormParams(r http.Request) (url.Values, error) {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			return nil, err
		}
	} else {
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
	}
	return r.Form, nil
}
