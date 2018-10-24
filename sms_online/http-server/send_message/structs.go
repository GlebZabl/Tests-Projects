package send_message

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

type Response struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
}

func (r *Response) send(res http.ResponseWriter) {
	QuData, err := json.Marshal(r)
	if err != nil {
		r.Status = false
		r.Message = err.Error()
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(QuData)
}
