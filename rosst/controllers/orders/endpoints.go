package orders

import (
	"Tests-Projects/rosst/envirement/DTOs/requests"
	"Tests-Projects/rosst/envirement/DTOs/responses"
	"Tests-Projects/rosst/envirement/services"
)

func MakeOrder(env *services.ServiceLocator, req requests.Request, response responses.Response) {
	logger := services.Logger()
	reqData := *req.(*requests.MakeOrder)
	defer func() {
		if err := response.Send(); err != nil {
			logger.Log(err)
		}
	}()

	userId, err := checkOrCreateUserByMail(env.IdStore(), env.UsersStore(), reqData.Mail)
	if err != nil {
		response.SetFailed()
		logger.Log(err)
		return
	}

	orderId, err := saveOrder(env.IdStore(), env.OrdersStore(), reqData.OrderInfo, userId)
	if err != nil {
		response.SetFailed()
		logger.Log(err)
		return
	}

	response.SetData(responses.MakeOrder{
		OrderId: orderId,
	})
	return
}

func GetOrders(env *services.ServiceLocator, req requests.Request, response responses.Response) {
	logger := services.Logger()
	reqData := *req.(*requests.GetOrders)
	defer func() {
		if err := response.Send(); err != nil {
			logger.Log(err)
		}
	}()

	user, err := getUserByMail(env.UsersStore(), reqData.Mail)
	if err != nil {
		response.SetFailed()
		logger.Log(err)
		return
	}

	if user == nil {
		response.SetError("this user have no orders")
		return
	}

	orders, err := getOrdersForUser(env.OrdersStore(), user.Id)
	if err != nil {
		response.SetFailed()
		logger.Log(err)
		return
	}

	response.SetData(responses.GetOrders{
		Orders: responses.ConvertToGetOrdersResponse(orders),
	})
}
