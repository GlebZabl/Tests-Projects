package structs

import (
	"time"
)

type timer struct {
	time      time.Duration
	age       int
	inChanel  chan bool
	outChanel chan bool
}

//запуск таймера
func (t *timer) Start() {
	selfChan := make(chan int)
	go t.Sleep(selfChan, t.age)
	for {
		select {
		case stoped := <-selfChan:
			if stoped == t.age {
				t.outChanel <- true
			}
		case <-t.inChanel:
			t.age++
			go t.Sleep(selfChan, t.age)
		}
	}
}

//по истечению времени возвращает номер эпохи в которой был запущен
func (t *timer) Sleep(deadLine chan int, age int) {
	time.Sleep(t.time)
	deadLine <- age
}
