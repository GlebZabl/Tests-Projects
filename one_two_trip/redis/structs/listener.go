package structs

import (
	. "Tests-Projects/one_two_trip/redis/constants"
	"Tests-Projects/one_two_trip/redis/functions"
	"fmt"
	. "github.com/AgileBits/go-redis-queue/redisqueue"
	"gopkg.in/redis.v2"
	"time"
)

type Listener struct {
	message      string
	fromTimer    chan bool
	toTimer      chan bool
	needToCheck  chan bool
	TasksQueue   *Queue
	ErrQueue     *Queue
	NotifyClient *redis.Client
}

//если не поступает новых сообщений переходим в режим выборов а если поступают то пытаемся взять из очереди
func (l *Listener) Work() {
	l.needToCheck = make(chan bool)
	l.fromTimer = make(chan bool)
	l.toTimer = make(chan bool)

	go l.listen()

	//получаем необработанные сообщения(если до этого работало только 1 приложение в режиме генератора и было прервано раньше чем актуальное включилось)
	l.tryGetMsg()

	timer := time.NewTimer(7 * time.Second)
	for {
		select {
		case <-l.needToCheck:
			timer = time.NewTimer(600 * time.Millisecond)
			l.tryGetMsg()
		case <-timer.C:
			return
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
			return
		}

		switch msg.(type) {
		case *redis.Message:
			l.needToCheck <- true
		}
	}
}

//проверяем очередь на наличие сообщений
func (l *Listener) tryGetMsg() {
	for {
		msg, err := l.TasksQueue.Pop()
		if err != nil || msg == "" {
			return
		}
		l.message = msg
		if !functions.GetError() {
			l.doSomethingWith(l.message)
		} else {
			l.pushError()
		}
	}
}

//возвращаем ошибку в redis
func (l *Listener) pushError() {
	l.ErrQueue.Push(l.message)
}

//так как в тз не сказано что делать с сообщениями будем выводить их в консоль
func (l *Listener) doSomethingWith(message string) {
	fmt.Println(message)
}
