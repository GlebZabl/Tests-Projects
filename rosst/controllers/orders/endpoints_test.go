package orders

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"Tests-Projects/rosst/controllers"
	"Tests-Projects/rosst/controllers/middlewares"
	"Tests-Projects/rosst/envirement/config"
	"Tests-Projects/rosst/envirement/models"
	"Tests-Projects/rosst/envirement/services"
	"Tests-Projects/rosst/envirement/services/logger"
	"Tests-Projects/rosst/envirement/services/ordersStore"
	"Tests-Projects/rosst/envirement/services/usersStore"
)

const mocGetOrdersResponse = `{"error":"","data":{"orders":[{"orderId":"testID2","orderInfo":"testID2","date":"testDate2"},{"orderId":"testID2","orderInfo":"testID2","date":"testDate2"}]}}`

func TestMakeOrder(t *testing.T) {

	_ = services.InitTestEnv(config.ConfOptions{
		LoggerKey: logger.ConsolePrinterName,
		IsTestEnv: true,
	})

	req, err := http.NewRequest("GET", controllers.MakeOrderUrl+"?mail=foomail.ru&orderInfo=bar", nil)
	if err != nil {
		t.Fail()
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Validation(MakeOrder))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fail()
		return
	}

	req, err = http.NewRequest("GET", controllers.MakeOrderUrl+"?mail=foo@mail.ru&orderInfo=bar", nil)
	if err != nil {
		t.Fatal()
		return
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middlewares.Validation(MakeOrder))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fail()
		return
	}

	uStore, ok := services.GetEnvironment().UsersStore().(*usersStore.TestStore)
	if !ok {
		t.Fail()
		return
	}

	oStore, ok := services.GetEnvironment().OrdersStore().(*ordersStore.TestStore)
	if !ok {
		t.Fail()
		return
	}

	failed := true
	for _, u := range uStore.GetData() {
		if u.Mail == "foo@mail.ru" {
			failed = false
		}
	}
	if failed {
		t.Fail()
		return
	}

	for _, o := range oStore.GetData() {
		if o.OwnerId == "mail=foo@mail.ru" && o.Info == "bar" {
			failed = false
		}
	}
	if failed {
		t.Fail()
		return
	}
}

func TestGetOrders(t *testing.T) {

	_ = services.InitTestEnv(config.ConfOptions{
		LoggerKey: logger.ConsolePrinterName,
		IsTestEnv: true,
	})

	req, err := http.NewRequest("GET", controllers.GetOrdersUrl+"?mail=foomail.ru&orderInfo=bar", nil)
	if err != nil {
		t.Fail()
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Validation(GetOrders))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fail()
		return
	}

	uStore, ok := services.GetEnvironment().UsersStore().(*usersStore.TestStore)
	if !ok {
		t.Fail()
		return
	}

	uStore.SetData(map[string]models.User{"testUserId": {
		Id:   "testUserId",
		Mail: "foo@mail.ru",
	}})

	oStore, ok := services.GetEnvironment().OrdersStore().(*ordersStore.TestStore)
	if !ok {
		t.Fail()
		return
	}

	oStore.SetData(map[string]models.Order{
		"testId": {
			Id:      "testID",
			OwnerId: "testUserId",
			Info:    "testInfo",
			Date:    "testDate",
		},
		"testId2": {
			Id:      "testID2",
			OwnerId: "testUserId",
			Info:    "testInfo2",
			Date:    "testDate2",
		},
	})

	req, err = http.NewRequest("GET", controllers.GetOrdersUrl+"?mail=foo@mail.ru", nil)
	if err != nil {
		t.Fatal()
		return
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middlewares.Validation(GetOrders))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fail()
		return
	}

	raw := rr.Body
	body, err := ioutil.ReadAll(raw)
	if err != nil {
		t.Fail()
		return
	}

	respData := string(body)
	if respData != mocGetOrdersResponse {
		t.Fail()
		return
	}
}
