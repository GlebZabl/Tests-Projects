package ordersStore

import "Tests-Projects/rosst/envirement/models"

type OrderStore interface {
	SaveOrder(o models.Order) error
	GetUsersOrders(userId string) ([]*models.Order, error)
}
