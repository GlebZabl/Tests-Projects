package responses

import "test_project/envirement/models"

type MakeOrder struct {
	OrderId string `json:"orderId"`
}

type GetOrders struct {
	Orders []order `json:"orders"`
}

type order struct {
	OrderId   string `json:"orderId"`
	OrderInfo string `json:"orderInfo"`
	Date      string `json:"date"`
}

func ConverToGetOrdersResponse(orders []*models.Order) []order {
	result := make([]order, 0, len(orders))
	for _, o := range orders {
		result = append(result, order{
			OrderId:   o.Id,
			OrderInfo: o.Id,
			Date:      o.Date,
		})
	}
	return result
}
