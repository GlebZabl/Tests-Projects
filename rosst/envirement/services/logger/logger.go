package logger

import (
	"fmt"
	"test_project/envirement/errors"
	"time"
)

const logFilePath = "./logs.txt"

type Logger interface {
	Log(source interface{})
}

type mainLogger struct {
	out printer
}

func (l *mainLogger) Log(source interface{}) {
	if source == nil {
		return
	}
	switch source.(type) {
	case *Errors.MainError:
		_ = l.out.print(time.Now().Format(time.RFC1123) + ": Error!" + "\n" + source.(*Errors.MainError).GetLogInfo())
	default:
		_ = l.out.print(time.Now().Format(time.RFC1123) + ": " + fmt.Sprint(source) + "\n")
	}
}

func (l *mainLogger) SetPrinter(printer printer) {
	l.out = printer
	l.out.setExternal(l)
}

func NewLogger(key string) Logger {
	printer := getLoggerInitializer(key)()
	logger := mainLogger{}
	logger.SetPrinter(printer)

	return &logger
}

type printer interface {
	print(string) error
	setExternal(*mainLogger)
}
