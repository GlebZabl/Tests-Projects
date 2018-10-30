package main

import (
	"fmt"
	"gopkg.in/redis.v2"
)

func main() {

	chanelClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
	})
	err := chanelClient.Ping()
	if err.Err() != nil {
		return
	}

	pubSub := chanelClient.PubSub()
	defer pubSub.Close()
	pubSub.Subscribe("l1")
	for {
		msg, err := pubSub.Receive()

		if err != nil {
			fmt.Println("error")
		}

		switch msg.(type) {
		case *redis.Message:
			fmt.Println(msg.(*redis.Message).Payload)
		}
	}
}


