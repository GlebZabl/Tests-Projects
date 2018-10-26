package message_saver

import (
	"encoding/json"
	"src/github.com/streadway/amqp"
)

//функция опрашивающая очередь
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

	for{
		msgs, err := chanel.Consume(
			QueueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil{
			reload<-true
			return
		}

		for msg := range msgs{
			data := new(Message)
			err = json.Unmarshal([]byte(string(msg.Body)),data)
			SaveMessage(*data)
		}
	}
}

//функция отвечающая за сохранение сообщение в бд или возврат его в очередь
func SaveMessage(msg Message)  {
		if  InsertMessage(msg) == nil{
			return
		}else{
			beckToQueue(msg)
		}
}

//функция отвечающая за возврат необработанного сообщения в очередь
func beckToQueue(msg Message){
	QuData := msg

	jsonData, err := json.Marshal(QuData)
	if err != nil{
		return
	}

	conn, err := amqp.Dial(QueueConnectionString)
	if err != nil{
		return
	}

	defer conn.Close()

	chanel,err := conn.Channel()
	if err!=nil {
		return
	}

	defer chanel.Close()

	err = chanel.Publish(
		"",
		QueueName,
		false,
		false,
		amqp.Publishing {
			ContentType: "application/json",
			Body:        jsonData,
		})
}