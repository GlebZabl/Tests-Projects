package middlewares

import (
	"Tests-Projects/rosst/controllers"
	"Tests-Projects/rosst/envirement/DTOs/requests"
	"Tests-Projects/rosst/envirement/errors"
	"reflect"
)

var registry = map[string]requests.Request{
	controllers.MakeOrderUrl: &requests.MakeOrder{},
	controllers.GetOrdersUrl: &requests.GetOrders{},
}

func GetRequest(url string) (requests.Request, error) {
	request, ok := registry[url]
	if !ok {
		return nil, Errors.NewWithMessage("wrong url path")
	}

	val := reflect.ValueOf(request)
	typ := val.Type().Elem()

	suragat := reflect.New(typ).Interface()
	result, ok := suragat.(requests.Request)
	if !ok {
		return nil, Errors.NewWithMessage("cant convert to request")
	}

	return result, nil
}
