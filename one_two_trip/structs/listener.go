package structs

import (
	. "Tests-Projects/one_two_trip/constants"
	. "Tests-Projects/one_two_trip/functions"
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
	startVote    chan bool
	TasksQueue   *Queue
	ErrQueue     *Queue
	NotifyClient *redis.Client
}

//если не поступает новых сообщений переходим в режим выборов а если поступают то пытаемся взять из очереди
func (l *Listener) Work() {
	l.needToCheck = make(chan bool)
	l.fromTimer = make(chan bool)
	l.toTimer = make(chan bool)
	l.startVote = make(chan bool)

	go l.listenVote()
	go l.listen()
	timer := time.NewTimer(4*time.Second)
	for {
		select {
		case <-l.needToCheck:
			timer = time.NewTimer(700*time.Millisecond)
			fmt.Println("get it")
			if l.tryGetMsg() {
				if !GetError() {
					fmt.Println(l.message)
				} else {
					l.pushError()
				}
			}
		case <-l.startVote:
			return
		case <-timer.C:
			fmt.Println("initialize vote")
			l.initializeVote()
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

//слушаем канал сообщающий о начале выборов
func (l *Listener) listenVote() {
	pubSub := l.NotifyClient.PubSub()
	defer pubSub.Close()
	pubSub.Subscribe(VoteNotifier)
	for {
		msg, err := pubSub.Receive()

		if err != nil {
			return
		}

		switch msg.(type) {
		case *redis.Message:
			l.startVote <- true
		}
	}
}

//проверяем очередь на наличие сообщений
func (l *Listener) tryGetMsg() bool {
	msg, err := l.TasksQueue.Pop()
	if err != nil || msg == "" {
		return false
	}
	l.message = msg
	return true
}

//возвращаем ошибку в redis
func (l *Listener) pushError() {
	l.ErrQueue.Push(l.message)
}

//оповещаем о начале выборов
func (l *Listener) initializeVote() {
	l.NotifyClient.Publish(VoteNotifier, "Vote!")
}
