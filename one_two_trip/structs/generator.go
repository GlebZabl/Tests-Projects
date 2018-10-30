package structs

import (
	. "Tests-Projects/one_two_trip/constants"
	"fmt"
	. "github.com/AgileBits/go-redis-queue/redisqueue"
	"gopkg.in/redis.v2"
	"strconv"
	"time"
)

type Generator struct {
	messageSanded int
	Name          string
	TasksQueue    *Queue
	message       string
	temp          int
	NotifyClient  *redis.Client
}

//отправляем рандомные сообщения раз в 500 мс
func (g *Generator) Work() {
	g.prepare()
	g.messageSanded = 0
	for {
		g.generateMessage()
		g.sendMessage()
		fmt.Println("send" + g.message)
		time.Sleep(500 * time.Millisecond)
	}
}

//пушим сообщение в очередь и публикуем в канал инфу о том что было опубликовано новое сообщение
func (g *Generator) sendMessage() {
	g.TasksQueue.Push(g.message)
	g.NotifyClient.Publish(ChanelName, "new message ready")
}

//как рандомное сообщение будем отправлять уникальное имя участника сети и номер сообщения среди отправленных этим участником(чтоб было нагляднее)
func (g *Generator) generateMessage() {
	g.message = g.Name + "  " + strconv.Itoa(g.messageSanded)
	g.messageSanded++
}

//чистим очередь перед началом
func (g *Generator) prepare() {
	for {
		msg, err := g.TasksQueue.Pop()

		if err != nil || msg == "" {
			return
		}

	}
}
