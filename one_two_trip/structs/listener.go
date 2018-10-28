package structs

import (
	. "Tests-Projects/one_two_trip/constants"
	"fmt"
	. "github.com/adeven/redismq"
	"gopkg.in/redis.v2"
)

type Listener struct {
	message string
	fromTimer chan bool
	toTimer chan bool
	needToCheck chan bool
	TasksQueue *Queue
	QueueConsumer *Consumer
	ErrQueue *Queue
	NotifyClient *redis.Client
}

//если не поступает новых сообщений переходим в режим выборов а если поступают то пытаемся взять из очереди
func (l *Listener) Work()  {
	l.needToCheck = make(chan bool)
	l.fromTimer = make(chan bool)
	l.toTimer = make(chan bool)

	go l.listen()

	pubSub := l.NotifyClient.PubSub()
	defer pubSub.Close()

	pubSub.Subscribe(ChanelName)
	for {
		msg, err := pubSub.Receive()

		if err != nil {
			fmt.Println("error")
		}

		switch msg.(type) {
		case *redis.Message:
			fmt.Println("get from chanel")
			fmt.Println(msg)
			l.needToCheck <- true
		}
	}
}

//слушаем канал сообщающий о появлении новых сообщений
func (l *Listener) listen()  {
	for{
		select {
		case <-l.fromTimer:
			return
		case <-l.needToCheck:
			/*l.toTimer <- true
			if l.tryGetMsg(){
				if !GetError() {
					fmt.Println(l.message)
				} else {
					l.pushError()
				}
			}*/
		}
	}
}

//проверяем очередь на наличие сообщений
func (l *Listener) tryGetMsg() bool {
	p, err := l.QueueConsumer.Get()
	if err != nil {
		return  true
	}
	p.Ack()
	l.message = p.Payload
	return false
}

//возвращаем ошибку в redis
func (l *Listener) pushError()  {
	l.ErrQueue.Put(l.message)
}

//переходим в режим генератора
func (l *Listener) becomeGenerator()  {
	generator := Generator{NotifyClient:l.NotifyClient,TasksQueue:l.TasksQueue}
	generator.Work()
}