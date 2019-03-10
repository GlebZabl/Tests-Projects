package ordersStore

import "test_project/envirement/models"

type OrderStore interface {
	SaveOrder(o models.Order) error
	GetUsersOrders(userId string) ([]*models.Order, error)
}
