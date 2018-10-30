package main

import (
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	"github.com/gomodule/redigo/redis"
	r "gopkg.in/redis.v2"
)

func main1() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
	}
	defer c.Close()

	q := redisqueue.New("tasksQueue", c)

	for {
		job, err := q.Pop()
		if err != nil {
			fmt.Println(job)
		}
		if job == ""{
			return
		}
	}
}

func main()  {
	chanelClient := r.NewClient(&r.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
	})
	err := chanelClient.Ping()
	if err.Err() != nil {
		return
	}
	chanelClient.Publish("l1","hi")
}