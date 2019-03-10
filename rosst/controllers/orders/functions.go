package orders

import (
	"time"

	"Tests-Projects/rosst/envirement/models"
	"Tests-Projects/rosst/envirement/services/idStore"
	"Tests-Projects/rosst/envirement/services/ordersStore"
	"Tests-Projects/rosst/envirement/services/usersStore"
)

func checkOrCreateUserByMail(iStore idStore.IdStore, uStore usersStore.UsersStore, mail string) (string, error) {
	user, err := getUserByMail(uStore, mail)
	if err != nil {
		return "", err
	}

	if user != nil {
		return user.Id, nil
	}

	user, err = createUser(iStore, uStore, mail)
	if err != nil {
		return "", err
	}

	return user.Id, nil
}

func getUserByMail(uStore usersStore.UsersStore, mail string) (*models.User, error) {
	user, err := uStore.GetUserByMail(mail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func createUser(iStore idStore.IdStore, uStore usersStore.UsersStore, mail string) (*models.User, error) {
	userID, err := iStore.GetNewId()
	user := models.User{
		Id:   userID,
		Mail: mail,
	}

	err = uStore.AddUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func saveOrder(iStore idStore.IdStore, oStore ordersStore.OrderStore, orderInfo string, userId string) (string, error) {
	orderID, err := iStore.GetNewId()
	if err != nil {
		return "", err
	}

	order := models.Order{
		Id:      orderID,
		OwnerId: userId,
		Info:    orderInfo,
		Date:    time.Now().Format("02 Jan 06 15:04"),
	}

	err = oStore.SaveOrder(order)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func getOrdersForUser(oStore ordersStore.OrderStore, userId string) ([]*models.Order, error) {
	orders, err := oStore.GetUsersOrders(userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
