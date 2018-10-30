package main

import (
	. "Tests-Projects/one_two_trip/constants"
	. "Tests-Projects/one_two_trip/functions"
	. "Tests-Projects/one_two_trip/structs"
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	rd "github.com/gomodule/redigo/redis"
	"gopkg.in/redis.v2"
	"strconv"
)

func main() {
	if GetMode() {
		ReadErrors()
	} else {
		//подключаемся к очередям
		con, err := rd.Dial("tcp",RedisConString+":"+RedisPort)
		if err != nil {
			return
		}
		defer con.Close()

		tasksQueue := redisqueue.New(TasksQueueName, con)
		errQueue := redisqueue.New(ErrQueueName, con)

		port, err := strconv.Atoi(RedisPort)
		if err != nil {
			return
		}

		//создаём клиент для прослушивания канала
		chanelClient := redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%d", RedisConString, port),
			DB:       RedisDbNumber,
		})

		cmdErr := chanelClient.Ping()
		if cmdErr.Err() != nil {
			return
		}

		client := Listener{TasksQueue:tasksQueue,ErrQueue: errQueue, NotifyClient: chanelClient}
		client.Work()
		fmt.Println("start vote")
	}
}
