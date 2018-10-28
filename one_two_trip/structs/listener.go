package structs

import (
	. "Tests-Projects/one_two_trip/constants"
	. "Tests-Projects/one_two_trip/functions"
	"fmt"
	. "github.com/adeven/redismq"
	"gopkg.in/redis.v2"
)

type Listener struct {
	message       string
	fromTimer     chan bool
	toTimer       chan bool
	needToCheck   chan bool
	TasksQueue    *Queue
	QueueConsumer *Consumer
	ErrQueue      *Queue
	NotifyClient  *redis.Client
}

//если не поступает новых сообщений переходим в режим выборов а если поступают то пытаемся взять из очереди
func (l *Listener) Work() {
	l.needToCheck = make(chan bool)
	l.fromTimer = make(chan bool)
	l.toTimer = make(chan bool)
	timer := timer{inChanel: l.toTimer, outChanel: l.fromTimer}

	go l.listen()
	go timer.Start()
	for {
		select {
		case <-l.needToCheck:
			if l.tryGetMsg() {
				if !GetError() {
					fmt.Println(l.message)
				} else {
					l.pushError()
				}
			}
		case <-l.fromTimer:
			println("start voute")
		}
	}
}

//слушаем канал сообщающий о появлении новых сообщений
func (l *Listener) listen() {
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
			l.toTimer <- true
			l.needToCheck <- true
		}
	}
}

//проверяем очередь на наличие сообщений
func (l *Listener) tryGetMsg() bool {
	p, err := l.QueueConsumer.Get()
	if err != nil {
		return false
	}
	p.Ack()
	l.message = p.Payload
	return true
}

//возвращаем ошибку в redis
func (l *Listener) pushError() {
	l.ErrQueue.Put(l.message)
}

//переходим в режим генератора
func (l *Listener) becomeGenerator() {
	generator := Generator{NotifyClient: l.NotifyClient, TasksQueue: l.TasksQueue}
	generator.Work()
}
