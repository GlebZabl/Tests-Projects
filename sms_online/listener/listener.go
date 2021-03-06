package main

import (
	. "Tests-Projects/sms_online/listener/message_saver"
)

func main() {
	reloadchan := make(chan bool)
	go Listen(reloadchan)

	//поднимаем заново если упало соединение с менеджером очередей
	for {
		select {
		case <-reloadchan:
			go Listen(reloadchan)
		}
	}
}
