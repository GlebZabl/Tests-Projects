package structs

import (
	"Tests-Projects/one_two_trip/constants"
	. "Tests-Projects/one_two_trip/functions"
	"github.com/adeven/redismq"
	"gopkg.in/redis.v2"
	"time"
)

type Generator struct {
	TasksQueue *redismq.Queue
	message string
	NotifyClient *redis.Client
}

//отправляем рандомные сообщения раз в 500 мс
func (g *Generator) Work() {
	for {
		g.generateMessage()
		g.sendMessage()
		time.Sleep(5 * time.Second)
	}
}

//генерируем рандомное сообщение
func (g *Generator) generateMessage() {
	g.message = GetRandomString()
}

//пушим сообщение в очередь и публикуем в канал инфу о том что было опубликовано новое сообщение
func (g *Generator) sendMessage() {
	g.TasksQueue.Put(g.message)
	g.NotifyClient.Publish(constants.ChanelName, "new message ready")
}

