package middlewares

import (
	"net/http"

	"Tests-Projects/rosst/envirement/DTOs/requests"
	"Tests-Projects/rosst/envirement/DTOs/responses"
	"Tests-Projects/rosst/envirement/services"
)

func Validation(next func(*services.ServiceLocator, requests.Request, responses.Response)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		logger := services.Logger()
		request, err := GetRequest(req.URL.Path)
		if err != nil {
			logger.Log(err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		err = Bind(request, *req)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if !request.Validate() {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		response := responses.NewResponse(res)
		env := services.GetEnvironment()

		next(env, request, response)
	}
}
