package main

import (
	. "Tests-Projects/one_two_trip/redis/constants"
	. "Tests-Projects/one_two_trip/redis/functions"
	. "Tests-Projects/one_two_trip/redis/structs"
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	rd "github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v2"
	"os"
)

func main() {
	if GetMode(os.Args) {
		ReadErrors()
	} else {
		//подключаемся к редису
		con, err := rd.Dial("tcp", RedisConString)
		if err != nil {
			return
		}
		defer con.Close()

		//работаем в режиме обработчика пока не перестанет поступать сигнал, затем проводим выборы и продолжаем работу в режиме установленном выборами
		for {
			success, message, listener, generator, voter := PrepareElements(con)
			if !success {
				println(message)
				return
			}

			fmt.Println("work as listener")
			listener.Work()
			fmt.Println("start vote")
			if voter.Vote() {
				fmt.Println("work as generator")
				generator.Work()
			}

			listener.NotifyClient.Close()
			generator.NotifyClient.Close()
			voter.PubClient.Close()
			voter.SubClient.Close()
		}
	}
}

//создаем элементы
func PrepareElements(con rd.Conn) (bool, string, *Listener, *Generator, *Voter) {

	//подключаемся к очередям
	tasksQueue := redisqueue.New(TasksQueueName, con)
	errQueue := redisqueue.New(ErrQueueName, con)
	//создаём клиент для прослушивания канала
	listenerChanelClient := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    RedisConString,
	})

	cmdErr := listenerChanelClient.Ping()
	if cmdErr.Err() != nil {
		return false,cmdErr.Err().Error(), nil, nil, nil
	}

	//создаём клиент для прослушивания канала в режиме голосования
	voteSubClient := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    RedisConString,
	})

	cmdErr = voteSubClient.Ping()
	if cmdErr.Err() != nil {
		return false, cmdErr.Err().Error(), nil, nil, nil
	}

	//создаём клиент для публикации в режиме голосования
	votePubClient := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    RedisConString,
	})

	cmdErr = votePubClient.Ping()
	if cmdErr.Err() != nil {
		return false, cmdErr.Err().Error(), nil, nil, nil
	}

	//получаем идентификатор для выборов
	uid, err := uuid.NewV4()
	if err != nil {
		return false, cmdErr.Err().Error(), nil, nil, nil
	}
	name := uid.String()

	listener := Listener{TasksQueue: tasksQueue, ErrQueue: errQueue, NotifyClient: listenerChanelClient}
	generator := Generator{TasksQueue: tasksQueue, NotifyClient: listenerChanelClient, Name: name}
	voter := Voter{Name: name, SubClient: voteSubClient, PubClient: votePubClient}

	return true, "", &listener, &generator, &voter
}
