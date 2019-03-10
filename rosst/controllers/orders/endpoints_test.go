package orders

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"test_project/controllers"
	"test_project/controllers/middlewares"
	"test_project/envirement/models"
	"test_project/envirement/services"
	"test_project/envirement/services/ordersStore"
	"test_project/envirement/services/usersStore"
	"testing"
)

const mocGetOrdersResponse = `{"error":"","data":{"orders":[{"orderId":"testID2","orderInfo":"testID2","date":"testDate2"},{"orderId":"testID2","orderInfo":"testID2","date":"testDate2"}]}}`

func TestMakeOrder(t *testing.T) {

	services.SetTestEnv()
	req, err := http.NewRequest("GET", controllers.MakeOrderUrl+"?mail=foomail.ru&orderInfo=bar", nil)
	if err != nil {
		t.Fail()
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Validation(MakeOrder))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fail()
	}

	req, err = http.NewRequest("GET", controllers.MakeOrderUrl+"?mail=foo@mail.ru&orderInfo=bar", nil)
	if err != nil {
		t.Fatal()
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middlewares.Validation(MakeOrder))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fail()
	}

	uStore, ok := services.GetEnvironment().UsersStore().(*usersStore.TestStore)
	if !ok {
		t.Fail()
	}

	oStore, ok := services.GetEnvironment().OrdersStore().(*ordersStore.TestStore)
	if !ok {
		t.Fail()
	}

	failed := true
	for _, u := range uStore.GetData() {
		if u.Mail == "foo@mail.ru" {
			failed = false
		}
	}
	if failed {
		t.Fail()
	}

	for _, o := range oStore.GetData() {
		if o.OwnerId == "mail=foo@mail.ru" && o.Info == "bar" {
			failed = false
		}
	}
	if failed {
		t.Fail()
	}
}

func TestGetOrders(t *testing.T) {

	services.SetTestEnv()
	req, err := http.NewRequest("GET", controllers.GetOrdersUrl+"?mail=foomail.ru&orderInfo=bar", nil)
	if err != nil {
		t.Fail()
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.Validation(GetOrders))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fail()
	}

	uStore, ok := services.GetEnvironment().UsersStore().(*usersStore.TestStore)
	if !ok {
		t.Fail()
	}

	uStore.SetData(map[string]models.User{"testUserId": {
		Id:   "testUserId",
		Mail: "foo@mail.ru",
	}})

	oStore, ok := services.GetEnvironment().OrdersStore().(*ordersStore.TestStore)
	if !ok {
		t.Fail()
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
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(middlewares.Validation(GetOrders))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fail()
	}

	raw := rr.Body
	body, err := ioutil.ReadAll(raw)
	if err != nil {
		t.Fail()
	}

	respData := string(body)
	if respData != mocGetOrdersResponse {
		t.Fail()
	}
}
