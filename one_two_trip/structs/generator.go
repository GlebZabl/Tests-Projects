package structs

import (
	. "Tests-Projects/one_two_trip/constants"
	. "Tests-Projects/one_two_trip/functions"
	"fmt"
	. "github.com/AgileBits/go-redis-queue/redisqueue"
	"gopkg.in/redis.v2"
	"strconv"
	"time"
)

type Generator struct {
	TasksQueue   *Queue
	message      string
	temp         int
	NotifyClient *redis.Client
}

//отправляем рандомные сообщения раз в 500 мс
func (g *Generator) Work() {
	i := 0
	for {
		//g.generateMessage()
		g.message = strconv.Itoa(i)
		i++
		g.sendMessage()
		fmt.Println("send" + g.message)
		time.Sleep(500 * time.Millisecond)
	}
}

//генерируем рандомное сообщение
func (g *Generator) generateMessage() {
	g.message = GetRandomString()
}

//пушим сообщение в очередь и публикуем в канал инфу о том что было опубликовано новое сообщение
func (g *Generator) sendMessage() {
	g.TasksQueue.Push(g.message)
	g.NotifyClient.Publish(ChanelName, "new message ready")
}
