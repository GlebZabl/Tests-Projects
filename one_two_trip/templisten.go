package main

import (
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	"github.com/gomodule/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
	}
	defer c.Close()

	q := redisqueue.New("some_queue_name", c)

	//wasAdded, err := q.Push("basic item")
	//if err != nil {
	//	fmt.Println(wasAdded)
	//}

	job, err := q.Pop()
	if err!=nil{
		fmt.Println(job)
	}

}