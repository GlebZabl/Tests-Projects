package main

import (
	. "Tests-Projects/sms_online/listener/message_saver"
)

func main(){
	reloadchan := make(chan bool)
	Listen(reloadchan)
}
