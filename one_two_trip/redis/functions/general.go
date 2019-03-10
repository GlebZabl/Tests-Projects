package functions

import (
	. "Tests-Projects/one_two_trip/redis/constants"
	"fmt"
	"github.com/AgileBits/go-redis-queue/redisqueue"
	rd "github.com/gomodule/redigo/redis"
	"math/rand"
)

//c 5ти процентной вероятностью возвращает ошибку
func GetError() bool {
	random := rand.Intn(19)
	return random == 0
}

//читаем очередь ошибок
func ReadErrors() {
	con, err := rd.Dial("tcp", RedisConString)
	if err != nil {
		return
	}
	defer con.Close()
	q := redisqueue.New(ErrQueueName, con)
	for {
		msg, err := q.Pop()

		if err != nil || msg == "" {
			return
		}
		fmt.Println(msg)
	}
}

//возвращает режим работы
func GetMode(args []string) bool {
	if len(args) > 1 {
		return args[1] == "getErrors"
	}
	return false
}
