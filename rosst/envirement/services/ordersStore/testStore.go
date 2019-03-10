package ordersStore

import (
	"Tests-Projects/rosst/envirement/errors"
	"Tests-Projects/rosst/envirement/models"
)

type TestStore struct{}

var (
	orders = map[string]models.Order{}
)

func init() {
	orders = make(map[string]models.Order)
}

func (t *TestStore) SaveOrder(o models.Order) error {
	if _, ok := orders[o.Id]; ok {
		return Errors.NewWithMessage("order with this Id is already exists")
	}
	orders[o.Id] = o
	return nil
}

func (t *TestStore) GetUsersOrders(userId string) ([]*models.Order, error) {
	var result []*models.Order
	for _, o := range orders {
		if o.OwnerId == userId {
			result = append(result, &o)
		}
	}
	return result, nil
}

func (t *TestStore) GetData() map[string]models.Order {
	return orders
}

func (t *TestStore) SetData(data map[string]models.Order) {
	orders = data
}

func NewTestStore() OrderStore {
	return new(TestStore)
}
