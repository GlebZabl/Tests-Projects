package structs

import (
	. "Tests-Projects/one_two_trip/constants"
	"gopkg.in/redis.v2"
	"time"
)

type Voter struct {
	Name       string
	SubClient  *redis.Client
	PubClient  *redis.Client
	selfChanel chan string
}

//начинаем слушать конкурентов, затем с небольшой задержкой(нужна чтоб все участники успели начали слушать) выдвигаем кандидатуру
func (v *Voter) Vote() bool {

	v.selfChanel = make(chan string)
	timer := time.NewTimer(400 * time.Millisecond)

	go v.listen()

	iAmTheKing := true
	time.Sleep(100 * time.Millisecond)
	v.PubClient.Publish(VoteChanel, v.Name)
	for {
		select {
		case candidate := <-v.selfChanel:
			if candidate > v.Name {
				iAmTheKing = false
			}
		case <-timer.C:
			return iAmTheKing
		}
	}
}

//слушаем кандидатов
func (v *Voter) listen() {
	pubSub := v.SubClient.PubSub()
	defer pubSub.Close()
	pubSub.Subscribe(VoteChanel)
	for {
		msg, err := pubSub.Receive()

		if err != nil {
			return
		}

		switch msg.(type) {
		case *redis.Message:
			v.selfChanel <- msg.(*redis.Message).Payload
		}
	}
}
