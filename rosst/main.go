package main

import (
	"net/http"

	"Tests-Projects/rosst/controllers"
	"Tests-Projects/rosst/controllers/middlewares"
	"Tests-Projects/rosst/controllers/orders"
	"Tests-Projects/rosst/envirement/config"
	"Tests-Projects/rosst/envirement/errors"
	"Tests-Projects/rosst/envirement/services"
)

const confPath = "./config.json"

func main() {
	for {
		if err := initialize(confPath); err != nil {
			services.Logger().Log("can`t initialize server components:")
			services.Logger().Log(err)
			break
		} else {
			services.Logger().Log("initialization is finished")
		}

		if err := serve(); err != nil {
			services.Logger().Log("cant start server")
			services.Logger().Log(err)
			break
		}
	}
}

func initialize(confPath string) error {
	cfg, err := config.LoadConfig(confPath)
	if err != nil {
		return err
	}

	err = services.InitServices(*cfg)
	if err != nil {
		return err
	}
	return nil
}

func createRouter() *http.ServeMux {
	result := http.NewServeMux()

	result.HandleFunc(controllers.MakeOrderUrl, middlewares.POST(middlewares.Validation(orders.MakeOrder)))
	result.HandleFunc(controllers.GetOrdersUrl, middlewares.GET(middlewares.Validation(orders.GetOrders)))

	return result
}

func serve() (err error) {
	logger := services.Logger()
	defer func() {
		if crash := recover(); crash != nil {
			err = Errors.NewWithMessage("Panic!")
		}
	}()

	server := &http.Server{Addr: services.Config().Port, Handler: createRouter()}

	errChan := make(chan error)
	go func(stopChan chan error) {
		err := server.ListenAndServe()
		if err != nil {
			stopChan <- err
		}
	}(errChan)

	logger.Log("server has started successfully and listening to port " + services.Config().Port[1:] + "!")
	for {
		select {
		case err := <-errChan:
			return err
		}
	}
}
