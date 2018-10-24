package main

import (
	"github.com/julienschmidt/httprouter"
	"Tests-Projects/sms_online/http-server/send_message"
	"net/http"
)

func main()  {
	router := httprouter.New()
	router.GET("/sendMessage/:message",send_message.GetMessage)
	http.ListenAndServe(":8000",router)
}
