package send_message

import (
	"fmt"
	"io/ioutil"
)

const (
	confPath = "http-server/config.txt"
)

var (
	QueueConnectionString string
	QueueName             string
	ErrMsg                string
)

func init() {
	QueueConnectionString = loadQueueConectionString()
	QueueName = "mainqueue"
	ErrMsg = "service is temporary unavailable"
}

func loadQueueConectionString() string {
	b, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
