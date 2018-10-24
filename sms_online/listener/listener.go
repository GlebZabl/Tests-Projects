package main

import (
	. "Tests-Projects/sms_online/listener/message_saver"
	"time"
)

func main(){
	message := Message{Text:"hi",GetByHTTPDate:time.Now().Unix(),GetByListenerDate:time.Now().Unix()}
	go SaveMessage(message)
	}
