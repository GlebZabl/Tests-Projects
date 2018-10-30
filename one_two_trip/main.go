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

		con, err := rd.Dial("tcp", RedisConString)
		if err != nil {
			return
		}
		defer con.Close()

		//подключаемся к очередям
		tasksQueue := redisqueue.New(TasksQueueName, con)
		errQueue := redisqueue.New(ErrQueueName, con)


		leader := false

		//работаем в режиме обработчика пока не перестанет поступать сигнал, затем проводим выборы и продолжаем работу в режиме установленном выборами
		for {
			//создаём клиент для прослушивания канала
			listenerChanelClient := redis.NewClient(&redis.Options{
				Network: "tcp",
				Addr:    RedisConString,
			})
			defer listenerChanelClient.Close()

			cmdErr := listenerChanelClient.Ping()
			if cmdErr.Err() != nil {
				return
			}

			//создаём клиент для прослушивания канала в режиме голосования
			voteSubClient := redis.NewClient(&redis.Options{
				Network: "tcp",
				Addr:    RedisConString,
			})
			defer voteSubClient.Close()

			cmdErr = voteSubClient.Ping()
			if cmdErr.Err() != nil {
				return
			}

			//создаём клиент для публикации в режиме голосования
			votePubClient := redis.NewClient(&redis.Options{
				Network: "tcp",
				Addr:    RedisConString,
			})
			defer votePubClient.Close()

			cmdErr = votePubClient.Ping()
			if cmdErr.Err() != nil {
				return
			}

			//получаем идентификатор для выборов
			uid, err := uuid.NewV4()
			if err != nil {
				return
			}
			name := uid.String()

			listener := Listener{TasksQueue: tasksQueue, ErrQueue: errQueue, NotifyClient: listenerChanelClient}
			generator := Generator{TasksQueue: tasksQueue, NotifyClient: listenerChanelClient, Name: name}
			voter := Voter{Name: name, SubClient: voteSubClient, PubClient: votePubClient}

			fmt.Println("started")
			if !leader {
				fmt.Println("work as listener")
				listener.Work()
				fmt.Println("start vote")
				leader = voter.Vote()
			} else {
				fmt.Println("work as generator")
				generator.Work()
			}
		}
	}
}
