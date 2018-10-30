package main

import (
	. "Tests-Projects/one_two_trip/constants"
	. "Tests-Projects/one_two_trip/functions"
	. "Tests-Projects/one_two_trip/structs"
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	rd "github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v2"
)

func main() {
	if GetMode() {
		ReadErrors()
	} else {

		//подключаемся к очередям
		con, err := rd.Dial("tcp",RedisConString)
		if err != nil {
			return
		}
		defer con.Close()

		tasksQueue := redisqueue.New(TasksQueueName, con)
		errQueue := redisqueue.New(ErrQueueName, con)


		//создаём клиент для прослушивания канала
		chanelClient := redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     RedisConString,
		})
		defer chanelClient.Close()

		cmdErr := chanelClient.Ping()
		if cmdErr.Err() != nil {
			return
		}

		uid,err:=uuid.NewV4()
		if err != nil{
			return
		}
		name := uid.String()

		listener := Listener{TasksQueue:tasksQueue,ErrQueue: errQueue, NotifyClient: chanelClient}
		generator := Generator{TasksQueue:tasksQueue,NotifyClient:chanelClient}
		voter := Voter{Name:name,Client:chanelClient}

		leader := false
		for {
			fmt.Println("started")
			if !leader {
				fmt.Println("work as listaner")
				listener.Work()
				fmt.Println("start vote")
				leader = voter.Vote()
			}else{
				fmt.Println("work as generator")
				generator.Work()
			}
		}
	}
}