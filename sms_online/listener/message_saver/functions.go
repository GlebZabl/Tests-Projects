package message_saver

import (
	"encoding/json"
	"fmt"
	"src/github.com/streadway/amqp"
)

func Listen(reload chan bool)  {
	conn, err := amqp.Dial(QueueConnectionString)
	if err != nil{
		reload <- true
		return
	}

	defer conn.Close()

	chanel,err := conn.Channel()
	if err!=nil {
		reload <- true
		return
	}

	defer chanel.Close()

	msgs, err := chanel.Consume(
		QueueName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for msg := range msgs{
		data := Message{}
		err = json.Unmarshal(msg.Body,data)
		SaveMessage(data)
	}
	fmt.Println(msgs)
}

func SaveMessage(msg Message)  {
		if  InsertMessage(msg) == nil{
			return
		}else{
			beckToQueue(msg)
		}
}

func beckToQueue(msg Message){
	return
}