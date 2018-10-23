package structs

import (
	"bufio"
	"os"
)

type FileParcer struct {
	parsers int
	stopChanel chan resultStruct
}

func (f *FileParcer) Parse(path string) (error,[][]string) {
	file, err := os.Open(path)
	if err != nil {
		return err,nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	f.parsers = 0
	f.stopChanel = make(chan resultStruct)

	for scanner.Scan() {
		f.parsers += 1
		lParser := lineParser{lineNumber:f.parsers-1,stopChanel:f.stopChanel,result:*new([]string)}
		go lParser.Parse(scanner.Text())
	}

	result := *new([][]string)
	for i:= 0; i < f.parsers; i++{
		result = append(result, *new([]string))
	}

	for {
		select {
		case lineParserStop := <-f.stopChanel:
			if lineParserStop.err{
				return *new(error), result
			}
			result[lineParserStop.lineNumber] = lineParserStop.data
			f.parsers -= 1
			if f.parsers == 0 {
				return nil, result
			}
		}
	}
}

type lineParser struct {
	lineNumber int
	stopChanel chan resultStruct
	result []string
}

func (l *lineParser) Parse(line string){
	err := false
	for i := 0;i<len(line);i++{
		if line[i:i+1] == "O" {
			l.result = append(l.result, "")
		}
		if line[i:i+1] == "X" {
			l.result = append(l.result, "X")
		}
		if line[i:i+1] != " " && line[i:i+1] != "X" && line[i:i+1] != "O"{
			err = true
		}
	}

	result := resultStruct{data:l.result, lineNumber:l.lineNumber,err:err}
	l.stopChanel <- result
}

type resultStruct struct {
	data []string
	lineNumber int
	err bool
}
