package responses

import (
	"Tests-Projects/rosst/envirement/errors"
	"encoding/json"
	"net/http"
)

type Response struct {
	writer http.ResponseWriter

	Error   string            `json:"error"`
	Data    interface{}       `json:"data"`
	headers map[string]string `json:"-"`
	status  bool              `json:"-"`
}

func (r *Response) Send() error {
	if !r.status {
		r.writer.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	jsonData, err := json.Marshal(r)
	if err != nil {
		return Errors.New(err)
	}

	for name, value := range r.headers {
		r.writer.Header().Set(name, value)
	}

	_, err = r.writer.Write(jsonData)
	if err != nil {
		return Errors.New(err)
	}

	return nil
}

func (r *Response) SetError(err string) {
	temp := err
	r.Error = temp
}

func (r *Response) SetFailed() {
	r.status = false
}

func (r *Response) SetData(data interface{}) {
	r.Data = data
}

func (r *Response) SetHeader(name string, value string) {
	r.headers[name] = value
}

func (r *Response) WriteHeader(code int) {
	r.writer.WriteHeader(code)
}

func NewResponse(w http.ResponseWriter) Response {
	return Response{
		Error: "",
		Data:  struct{}{},
		headers: map[string]string{
			"Content-Type": "application/json",
		},
		status: true,
		writer: w,
	}
}
