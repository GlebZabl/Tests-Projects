package message_saver

import (
	"fmt"
	"io/ioutil"
)

const(
	confPath = "listener/config.txt"
)

var(
	DBConnectionString string
	QueueConnectionString = "amqp://zabl:1334216@192.168.56.101:5672/"
	QueueName = "mainqueue"
)

func init() {
	DBConnectionString = loadDbConectionString(confPath)
}

func loadDbConectionString(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
