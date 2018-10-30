package structs

import (
	. "Tests-Projects/one_two_trip/constants"
	"gopkg.in/redis.v2"
	"time"
)

type Voter struct {
	Name string
	SubClient *redis.Client
	PubClient *redis.Client
	selfChanel chan string
}

func (v *Voter) Vote() bool  {

	v.selfChanel = make(chan string)
	timer := time.NewTimer(500*time.Millisecond)

	go v.listen()

	iAmTheKing := true
	v.PubClient.Publish(VoteChanel,v.Name)
	for{
		select {
		case condidate := <- v.selfChanel:
			if condidate>v.Name{
				iAmTheKing = false
			}
		case <- timer.C:
			return iAmTheKing
		}
	}
}

func (v *Voter) listen()  {
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
			v.PubClient.Publish(VoteChanel,v.Name)
			v.selfChanel<-msg.(*redis.Message).Payload
		}
	}
}
