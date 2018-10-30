package structs

import (
	."Tests-Projects/one_two_trip/constants"
	"gopkg.in/redis.v2"
	"time"
)

type Voter struct {
	Name string
	Client *redis.Client
}

func (v *Voter) Vote() bool  {

	fromListener := make(chan string)
	timer := time.NewTimer(5* time.Second)

	go v.listen(fromListener)

	iAmTheKing := true
	for{
		select {
		case condidate := <- fromListener:
			if condidate>v.Name{
				iAmTheKing = false
			}
		case <- timer.C:
			return iAmTheKing
		}
	}
}

func (v *Voter) listen(outChan chan string)  {
	pubSub := v.Client.PubSub()
	defer pubSub.Close()
	pubSub.Subscribe(VoteChanel)
	for {
		msg, err := pubSub.Receive()

		if err != nil {
			return
		}

		switch msg.(type) {
		case *redis.Message:
			outChan<-msg.(*redis.Message).Payload
		}
	}
}
