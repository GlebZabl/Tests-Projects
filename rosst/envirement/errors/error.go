package Errors

import (
	"runtime"
	"strconv"
)

type baseError struct {
	msg string
}

func (b *baseError) Error() string {
	return b.msg
}

type MainError struct {
	trace  []string
	origin error
}

func (m *MainError) Error() string {
	return m.origin.Error()
}

func (m *MainError) GetLogInfo() string {
	var result, spaces string
	for i := len(m.trace) - 1; i >= 0; i-- {
		result += spaces + m.trace[i] + "\n"
		spaces += "  "
	}
	result = result[:len(result)-2] + ":\n" + spaces + m.origin.Error()
	return result
}

func New(origin error) error {
	return &MainError{
		trace:  GetStackTrace(),
		origin: origin,
	}
}

func NewWithMessage(message string) error {
	return &MainError{
		trace:  GetStackTrace(),
		origin: &baseError{msg: message},
	}
}

func GetStackTrace() []string {
	var result []string
	lvl := 2
	for {
		pc, _, line, _ := runtime.Caller(lvl)
		file, line := runtime.FuncForPC(pc).FileLine(pc)
		funcName := runtime.FuncForPC(pc).Name()
		result = append(result, file+":"+strconv.Itoa(line))
		if funcName == "main.main" {
			return result
		}
		if funcName == "net/http.HandlerFunc.ServeHTTP" {
			return result[:len(result)-1]
		}
		lvl++
	}
}
