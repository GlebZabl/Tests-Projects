package main

import (
	"fmt"
	"github.com/adeven/redismq"
	"github.com/satori/go.uuid"
)


func main() {
	random,_ := uuid.NewV4()
	consume("listener"+random.String())
}

func consume(name string) {
	testQueue := redismq.CreateQueue("localhost", "6379", "", 1, "tasks")
	consumer, err := testQueue.AddConsumer(name)
	if err != nil {
		panic(err)
	}
	for {
		p, err := consumer.Get()
		if err != nil {
			continue
		}
		err = p.Ack()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(p.Payload)
	}
}
