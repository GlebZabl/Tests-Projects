package structs

import (
	"bufio"
	"os"
)

type chanSet struct {
	Highter     chan int
	Self        chan int
	Lower       chan int
	StopHighter chan bool
	StopSelf    chan bool
	StopLower   chan bool
}

type fullSet struct {
	from chanSet
	to   chanSet
}

var (
	chanelsMap map[int]*fullSet
)

type stopChanStruct struct {
	lineNumber int
	data       []int
	err        bool
}

type FileParser struct {
	scanner    *bufio.Scanner
	parsers    int
	stopChanel chan stopChanStruct
	result     [][]int
	FieldSize  int
}

func (f *FileParser) Parse(path string) (error, [][]int) {
	chanelsMap = make(map[int]*fullSet)

	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	f.scanner = bufio.NewScanner(file)
	f.parsers = 0
	f.stopChanel = make(chan stopChanStruct)

	return f.startScaning()
}

//запускает сканирование
func (f *FileParser) startScaning() (error, [][]int) {
	for f.scanner.Scan() {
		f.prepareChans()
		parser := lineParser{fieldSize: f.FieldSize, lineNumber: f.parsers, result: *new([]int), geted: *new([]int), resultChan: f.stopChanel, statuses: []bool{false, false, false}, inputChans: &chanelsMap[f.parsers].from, outputChans: &chanelsMap[f.parsers].to}
		for i := 0;i<=f.FieldSize;i++{
			parser.geted =  append(parser.geted, 0)
		}
		f.parsers += 1
		go parser.listen()
		go parser.listenReady()
		go parser.parse(f.scanner.Text())
	}
	for i := 0; i < f.parsers; i++ {
		f.result = append(f.result, []int{})
	}
	for {
		select {
		case end := <-f.stopChanel:
			if end.err {
				return *new(error), nil
			} else {
				f.parsers -= 1
				f.result[end.lineNumber] = end.data
				if f.parsers == 0 {
					return nil, f.result
				}
			}
		}
	}
}

func (f *FileParser) prepareChans() {
	self := make(chan int)
	stopself := make(chan bool)

	if f.parsers == 0 {
		temp := fullSet{}
		temp.to = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool)}
		temp.from = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool)}
		chanelsMap[f.parsers] = &temp
		return
	}

	if f.parsers == f.FieldSize {
		temp := fullSet{}
		temp.to = chanSet{Self: self, StopSelf: stopself, Highter: chanelsMap[f.parsers-1].from.Lower, StopHighter: chanelsMap[f.parsers-1].from.StopLower}
		temp.from = chanSet{Self: self, StopSelf: stopself, Highter: chanelsMap[f.parsers-1].to.Lower, StopHighter: chanelsMap[f.parsers-1].to.StopLower}
		chanelsMap[f.parsers] = &temp
		return
	}

	temp := fullSet{}
	temp.to = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool), Highter: chanelsMap[f.parsers-1].from.Lower, StopHighter: chanelsMap[f.parsers-1].from.StopLower}
	temp.from = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool), Highter: chanelsMap[f.parsers-1].to.Lower, StopHighter: chanelsMap[f.parsers-1].to.StopLower}
	chanelsMap[f.parsers] = &temp
}

//обработчик строки
type lineParser struct {
	fieldSize   int
	lineNumber  int
	result      []int
	geted       []int
	outputChans *chanSet
	inputChans  *chanSet
	resultChan  chan stopChanStruct
	statuses    []bool
}

//обрабатываем строчку
func (l *lineParser) parse(line string) {
	err := false
	for i := 0; i < len(line); i++ {
		if line[i:i+1] == "O" {
			l.result = append(l.result, 0)
		}
		if line[i:i+1] == "X" {
			l.result = append(l.result, -1)
			l.sendSignals(len(l.result)-1)
		}
		if line[i:i+1] != " " && line[i:i+1] != "X" && line[i:i+1] != "O" {
			err = true
		}
	}

	result := stopChanStruct{data: l.result, lineNumber: l.lineNumber, err: err}
	if err {
		l.resultChan <- result
		return
	}
	l.sendFinish()
}

//отправляем сообщения о минах обработчикам соседних строчек
func (l *lineParser) sendSignals(index int) {
	if l.lineNumber != 0 {
		for i := index - 1; i <= index+1; i++ {
			l.outputChans.Highter <- i
		}
	}
	if l.lineNumber != l.fieldSize {
		for i := index - 1; i <= index+1; i++ {
			l.outputChans.Lower <- i
		}
	}
	l.outputChans.Self <- index - 1
	l.outputChans.Self <- index + 1
}

//отправляем сообщения о том что завершили работу обработчикам соседних строчек
func (l *lineParser) sendFinish() {
	if l.lineNumber != 0 {
		l.outputChans.StopHighter <- true
	}
	if l.lineNumber != l.fieldSize {
		l.outputChans.StopLower <- true
	}
	l.outputChans.StopSelf <- true
}

//слушаем обработчики соседних строчек о минах
func (l *lineParser) listen() {
	for {
		select {
		case index := <-l.inputChans.Highter:
			if index != -1 && index <= l.fieldSize {
				l.geted[index] += 1
			}

		case index := <-l.inputChans.Self:
			if index != -1 && index <= l.fieldSize {
				l.geted[index] += 1
			}

		case index := <-l.inputChans.Lower:
			if index != -1 && index <= l.fieldSize {
				l.geted[index] += 1
			}
		}
	}
}

//слушаем обработчики соседних строчек о том что они закончили обработку
func (l *lineParser) listenReady() {
	for {
		select {
		case <-l.inputChans.StopHighter:
			l.statuses[0] = true
			if l.getStatus() {
				l.resultChan <- stopChanStruct{data: l.finalCount(), err: false, lineNumber: l.lineNumber}
			}
		case <-l.inputChans.StopSelf:
			l.statuses[1] = true
			if l.getStatus() {
				l.resultChan <- stopChanStruct{data: l.finalCount(), err: false, lineNumber: l.lineNumber}
			}
		case <-l.inputChans.StopLower:
			l.statuses[2] = true
			if l.getStatus() {
				l.resultChan <- stopChanStruct{data: l.finalCount(), err: false, lineNumber: l.lineNumber}
			}
		}
	}
}

//получаем статус(можно ли закрыть поток обрабатывающий строку)
func (l *lineParser) getStatus() bool {
	if l.lineNumber == 0 {
		return l.statuses[1] && l.statuses[2]
	}
	if l.lineNumber == l.fieldSize {
		return l.statuses[0] && l.statuses[1]
	}
	return l.statuses[0] && l.statuses[1] && l.statuses[2]
}

//соединяем полученные значения и мины
func (l *lineParser) finalCount() []int {
	result := *new([]int)
	for i := 0; i <= l.fieldSize; i++ {
		if l.result[i] == -1 {
			result = append(result, l.result[i])
		} else {
			result = append(result, l.geted[i])
		}
	}
	return result
}
