package send_message

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text          string `json:"text"`
	GetByHTTPDate int64  `json:"get_date"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

//ф-ия отправки ответа клиенту
func (r *Response) send(res http.ResponseWriter) {
	QuData, err := json.Marshal(r)
	if err != nil {
		r.Status = false
		r.Message = err.Error()
	}
	res.Header().Set("Content-Type", "application/json")
	if r.Status {
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(http.StatusBadGateway)
	}
	res.Write(QuData)
}
