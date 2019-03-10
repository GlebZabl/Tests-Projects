package ordersStore

import (
	"sync"

	"Tests-Projects/rosst/envirement/errors"
	"Tests-Projects/rosst/envirement/models"
)

type TestStore struct{}

var orders safeMap

type safeMap struct {
	sync.Mutex
	data map[string]models.Order
}

func init() {
	orders = safeMap{data: make(map[string]models.Order)}
}

func NewTestStore() OrderStore {
	return new(TestStore)
}

func (t *TestStore) SaveOrder(o models.Order) error {
	orders.Lock()
	defer orders.Unlock()

	if _, ok := orders.data[o.Id]; ok {
		return Errors.NewWithMessage("order with this Id is already exists")
	}
	orders.data[o.Id] = o
	return nil
}

func (t *TestStore) GetUsersOrders(userId string) ([]*models.Order, error) {
	orders.Lock()
	defer orders.Unlock()

	var result []*models.Order
	for _, o := range orders.data {
		if o.OwnerId == userId {
			result = append(result, &models.Order{
				Id:      o.Id,
				OwnerId: o.OwnerId,
				Info:    o.Info,
				Date:    o.Date,
			})
		}
	}
	return result, nil
}

//for tests only
func (t *TestStore) GetData() map[string]models.Order {
	orders.Lock()
	defer orders.Unlock()

	return orders.data
}

func (t *TestStore) SetData(data map[string]models.Order) {
	orders.Lock()
	defer orders.Unlock()

	orders.data = data
}
