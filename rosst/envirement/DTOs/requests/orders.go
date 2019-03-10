package requests

const mailRegExp = `[a-zA-Z]*@[a-zA-Z]*\.(ru|com|org|mail)`

type MakeOrder struct {
	Mail      string `query:"mail"`
	OrderInfo string `query:"orderInfo"`
}

func (m *MakeOrder) Validate() bool {
	if !checkRegExp(m.Mail, mailRegExp) {
		return false
	}
	if len(m.OrderInfo) > 100 || len(m.OrderInfo) < 3 {
		return false
	}

	return true
}

type GetOrders struct {
	Mail string `query:"mail"`
}

func (g *GetOrders) Validate() bool {
	if !checkRegExp(g.Mail, mailRegExp) {
		return false
	}
	return true
}
