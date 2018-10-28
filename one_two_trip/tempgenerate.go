package main

import (
	"fmt"
	"github.com/adeven/redismq"
	"time"
)

func main()  {
	fmt.Println("started")
	testQueue := redismq.CreateQueue("localhost", "6379", "", 9, "clicks")
	for {
		testQueue.Put("testpayload")
		time.Sleep(2*time.Second)
	}
}

