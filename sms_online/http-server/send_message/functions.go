package send_message

import (
	"encoding/json"
	"src/github.com/streadway/amqp"
	"time"
)

func AddToQueue(msg string) (error,string) {
	QuData := Message{Text:msg,GetByHTTPDate:time.Now().Unix()}

	jsonData, err := json.Marshal(QuData)
	if err != nil{
		return err,ErrMsg
	}

	conn, err := amqp.Dial(QueueConnectionString)
	if err != nil{
		return err,ErrMsg
	}

	defer conn.Close()

	chanel,err := conn.Channel()
	if err!=nil {
		return err, ErrMsg
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
	if err!= nil{
		return err,ErrMsg
	}

	return nil,""
}
