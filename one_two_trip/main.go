package main

import (
	. "Tests-Projects/one_two_trip/constants"
	. "Tests-Projects/one_two_trip/functions"
	. "Tests-Projects/one_two_trip/structs"
	"fmt"
	"github.com/adeven/redismq"
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v2"
	"strconv"
)

func main() {
	if GetMode() {
		ReadErrors()
	} else {
		//подключаемся к очередям
		tasksQueue := redismq.CreateQueue(RedisConString, RedisPort, RedisPassword, RedisDbNumber, TasksQueueName)
		errQueue := redismq.CreateQueue(RedisConString, RedisPort, RedisPassword, RedisDbNumber, ErrQueueName)

		name, err := uuid.NewV4()
		cnsmr, err := tasksQueue.AddConsumer(name.String())
		if err != nil {
			return
		}

		port, err := strconv.Atoi(RedisPort)
		if err != nil {
			return
		}

		//создаём клиент для прослушивания канала
		chanelClient := redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%d", RedisConString, port),
			Password: RedisPassword,
			DB:       RedisDbNumber,
		})

		cmdErr := chanelClient.Ping()
		if cmdErr.Err() != nil {
			return
		}

		client := Listener{QueueConsumer: cnsmr, ErrQueue: errQueue, NotifyClient: chanelClient, TasksQueue: tasksQueue}
		client.Work()
	}
}
