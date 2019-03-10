package responses

import "Tests-Projects/rosst/envirement/models"

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

func ConvertToGetOrdersResponse(orders []*models.Order) []order {
	result := make([]order, 0, len(orders))
	for i := range orders {
		result = append(result, order{
			OrderId:   orders[i].Id,
			OrderInfo: orders[i].Info,
			Date:      orders[i].Date,
		})
	}
	return result
}
