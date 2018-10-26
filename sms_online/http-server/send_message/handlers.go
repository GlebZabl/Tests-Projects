package send_message

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//хэндлер запроса
func GetMessage(res http.ResponseWriter, _ *http.Request, params httprouter.Params)  {
	message:= params.ByName("message")
	err,message := AddToQueue(message)
	response := Response{Status:err==nil,Message:message}
	response.send(res)
}

