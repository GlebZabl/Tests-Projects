package orders

import (
	"testing"

	"Tests-Projects/rosst/envirement/models"
	"Tests-Projects/rosst/envirement/services/idStore"
	"Tests-Projects/rosst/envirement/services/ordersStore"
	"Tests-Projects/rosst/envirement/services/usersStore"
)

func TestCheckOrCreateUserByMailFunc(t *testing.T) {
	iStore := idStore.NewTestStore()
	uStore := usersStore.NewTestStore()
	userId, _ := checkOrCreateUserByMail(iStore, uStore, "testUser")

	dataChanger, ok := uStore.(*usersStore.TestStore)
	if !ok {
		t.Fail()
	}

	user, ok := dataChanger.GetData()[userId]
	if !ok || user.Id != userId || user.Mail != "testUser" {
		t.Fail()
	}
}

func TestCheckUserByMailFunc(t *testing.T) {
	uStore := usersStore.NewTestStore()

	dataChanger, ok := uStore.(*usersStore.TestStore)
	if !ok {
		t.Fail()
	}
	dataChanger.SetData(make(map[string]models.User))

	user, _ := getUserByMail(uStore, "testMail")
	if user != nil {
		t.Fail()
	}

	dataChanger.SetData(map[string]models.User{"testId": {
		Id:   "testId",
		Mail: "testMail"}})

	user, _ = getUserByMail(uStore, "testMail")
	if user == nil || user.Id != "testId" || user.Mail != "testMail" {
		t.Fail()
	}
}

func TestCreateUserByMailFunc(t *testing.T) {
	iStore := idStore.NewTestStore()
	uStore := usersStore.NewTestStore()
	user, _ := createUser(iStore, uStore, "testUser")
	if user == nil {
		t.Fail()
		return
	}

	dataChanger, ok := uStore.(*usersStore.TestStore)
	if !ok {
		t.Fail()
		return
	}

	u, ok := dataChanger.GetData()[user.Id]
	if !ok || u.Id != user.Id || u.Mail != "testUser" {
		t.Fail()
		return
	}
}

func TestSaveOrderFunc(t *testing.T) {
	iStore := idStore.NewTestStore()
	oStore := ordersStore.NewTestStore()
	orderID, _ := saveOrder(iStore, oStore, "testInfo", "testId")

	dataChanger, ok := oStore.(*ordersStore.TestStore)
	if !ok {
		t.Fail()
		return
	}

	order, ok := dataChanger.GetData()[orderID]
	if !ok || order.Id != orderID || order.Info != "testInfo" || order.OwnerId != "testId" {
		t.Fail()
		return
	}
}

func TestGetOrdersFunc(t *testing.T) {
	oStore := ordersStore.NewTestStore()

	dataChanger, ok := oStore.(*ordersStore.TestStore)
	if !ok {
		t.Fail()
		return
	}
	dataChanger.SetData(make(map[string]models.Order))

	orders, _ := getOrdersForUser(oStore, "testUser")
	if len(orders) != 0 {
		t.Fail()
		return
	}

	dataChanger.SetData(map[string]models.Order{
		"testId": {
			Id:      "testId",
			OwnerId: "testUser",
			Info:    "testInfo",
			Date:    "testDate",
		},
	})

	orders, _ = getOrdersForUser(oStore, "testUser")
	if len(orders) != 1 || orders[0].Id != "testId" || orders[0].OwnerId != "testUser" ||
		orders[0].Info != "testInfo" || orders[0].Date != "testDate" {
		t.Fail()
		return
	}
}
