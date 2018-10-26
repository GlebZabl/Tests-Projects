package structs

import (
	"bufio"
	"os"
)

type chanSet struct {
	Higher     chan int
	Self       chan int
	Lower      chan int
	StopHigher chan bool
	StopSelf   chan bool
	StopLower  chan bool
}

type fullSet struct {
	from chanSet //каналы для отправки
	to   chanSet //каналы для получения
}

var (
	channelsMap map[int]*fullSet
)
//структура которую возвращает lineparser
type stopChanStruct struct {
	lineNumber int
	data       []int
	err        bool
}

//структура для обработки файла
type FileParser struct {
	scanner    *bufio.Scanner
	parsers    int                 //кол-во работающих обработчиков строк
	stopChanel chan stopChanStruct //канал для получения результата от обработчиков строк
	result     [][]int             //результат
	FieldSize  int                 //размер одной стороны поля
}

//сканирует и обсчитывает файл
func (f *FileParser) Parse(path string) (bool, [][]int) {
	channelsMap = make(map[int]*fullSet)

	file, err := os.Open(path)
	if err != nil {
		return true, nil
	}
	defer file.Close()

	f.scanner = bufio.NewScanner(file)
	f.parsers = 0
	f.stopChanel = make(chan stopChanStruct)

	return f.startScanning()
}

//запускает сканирование
func (f *FileParser) startScanning() (bool, [][]int) {
	for f.scanner.Scan() {
		f.prepareChans()
		parser := lineParser{fieldSize: f.FieldSize, lineNumber: f.parsers, result: *new([]int), geted: *new([]int), resultChan: f.stopChanel, statuses: []bool{false, false, false}, inputChans: &channelsMap[f.parsers].from, outputChans: &channelsMap[f.parsers].to}

		for i := 0; i <= f.FieldSize; i++ {
			parser.geted = append(parser.geted, 0)
		}

		f.parsers += 1
		go parser.listen()
		go parser.listenReady()
		go parser.parse(f.scanner.Text())
	}
	if f.parsers != f.FieldSize+1 {
		return true, nil
	}

	for i := 0; i < f.parsers; i++ {
		f.result = append(f.result, []int{})
	}

	return f.waitScanning()
}

//ждём пока завершат работу все обработчики
func (f *FileParser) waitScanning() (bool, [][]int) {
	for {
		select {
		case end := <-f.stopChanel:
			if end.err {
				return true, nil
			} else {
				f.parsers -= 1
				f.result[end.lineNumber] = end.data
				if f.parsers == 0 {
					return false, f.result
				}
			}
		}
	}
}

//готовим элемент карты для обработчика
func (f *FileParser) prepareChans() {
	self := make(chan int)
	stopself := make(chan bool)

	if f.parsers == 0 {
		temp := fullSet{}
		temp.to = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool)}
		temp.from = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool)}
		channelsMap[f.parsers] = &temp
		return
	}

	if f.parsers == f.FieldSize {
		temp := fullSet{}
		temp.to = chanSet{Self: self, StopSelf: stopself, Higher: channelsMap[f.parsers-1].from.Lower, StopHigher: channelsMap[f.parsers-1].from.StopLower}
		temp.from = chanSet{Self: self, StopSelf: stopself, Higher: channelsMap[f.parsers-1].to.Lower, StopHigher: channelsMap[f.parsers-1].to.StopLower}
		channelsMap[f.parsers] = &temp
		return
	}

	temp := fullSet{}
	temp.to = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool), Higher: channelsMap[f.parsers-1].from.Lower, StopHigher: channelsMap[f.parsers-1].from.StopLower}
	temp.from = chanSet{Self: self, StopSelf: stopself, Lower: make(chan int), StopLower: make(chan bool), Higher: channelsMap[f.parsers-1].to.Lower, StopHigher: channelsMap[f.parsers-1].to.StopLower}
	channelsMap[f.parsers] = &temp
}

//обработчик строки
type lineParser struct {
	fieldSize   int                 //размер поля(одной стороны)
	lineNumber  int                 //номер обрабатываемой строки
	result      []int               //результат обработки строки
	geted       []int               //информация о бомбах от обработчиков соседних строк
	outputChans *chanSet            //каналы для откправки сигналов
	inputChans  *chanSet            //каналы для прослушивания сигналов
	resultChan  chan stopChanStruct //канал для отправки результата
	statuses    []bool              //статусы завершения обработчиков(этого и двух соседних)
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
			l.sendSignals(len(l.result) - 1)
		}
		if line[i:i+1] != " " && line[i:i+1] != "X" && line[i:i+1] != "O" {
			err = true
		}
	}

	if len(l.result) != l.fieldSize+1 {
		err = true
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
			l.outputChans.Higher <- i
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
		l.outputChans.StopHigher <- true
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
		case index := <-l.inputChans.Higher:
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
		case <-l.inputChans.StopHigher:
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
