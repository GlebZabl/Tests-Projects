package send_message

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetMessage(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	message:= params.ByName("message")
	err,message := AddToQueue(message)
	response := Response{Status:err==nil,Message:message}
	response.send(res)
}

