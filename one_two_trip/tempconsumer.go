package main

import (
	"fmt"
	"gopkg.in/redis.v2"
)

func main() {

	chanelClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
		Password: "",
		DB:       6,
	})
	err := chanelClient.Ping()
	if err.Err() != nil {
		return
	}

	firstc := make(chan string)
	secondc := make(chan string)
	go first(chanelClient,firstc)
	go second(chanelClient,secondc)
	chanelClient.Publish("testchanel1","self")
	for{
		select {
		case a:=<-firstc:
			println("get from first: " + a)
		case a:=<-secondc:
			println("get from second: " + a)
		}
	}
}

func first(chanelClient *redis.Client, out chan string) {
	pubSub := chanelClient.PubSub()
	defer pubSub.Close()
	pubSub.Subscribe("infoChanel")
	for {
		msg, err := pubSub.Receive()

		if err != nil {
			fmt.Println("error")
		}

		switch msg.(type) {
		case *redis.Message:
			out <- msg.(*redis.Message).Payload
		}
	}
}

func second(chanelClient *redis.Client, out chan string) {
	pubSub := chanelClient.PubSub()
	defer pubSub.Close()
	pubSub.Subscribe("testchanel2")
	for {
		msg, err := pubSub.Receive()

		if err != nil {
			fmt.Println("error")
		}

		switch msg.(type) {
		case *redis.Message:
			out <- msg.(*redis.Message).Payload
		}
	}
}
