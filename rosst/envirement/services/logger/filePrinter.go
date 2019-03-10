package logger

import (
	"os"

	"Tests-Projects/rosst/envirement/errors"
)

type filePrinter struct {
	external *mainLogger
	file     logFile
	channel  chan string
}

func (f *filePrinter) setExternal(logger *mainLogger) {
	f.external = logger
}

func (f *filePrinter) print(info string) error {
	f.channel <- info
	return nil
}

func (f *filePrinter) listen() {
	defer func() {
		if err := recover(); err != nil {
			printer := newConsolePrinter()
			f.external.SetPrinter(printer)
			f.external.Log("file mainLogger has crashed with panic:")
			f.external.Log(err)
		}
	}()

	for {
		select {
		case info := <-f.channel:
			err := f.file.write([]byte(info))

			if err != nil {
				printer := newConsolePrinter()
				f.external.SetPrinter(printer)
				f.external.Log("file mainLogger is crashed:")
				f.external.Log(err)
			}
		}
	}
}

func newFilePrinter() printer {
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return newConsolePrinter()
	}

	result := filePrinter{file: logFile{out: file}, channel: make(chan string, 100)}
	go result.listen()
	return &result
}

type logFile struct {
	out *os.File
}

func (l *logFile) write(info []byte) error {
	_, err := l.out.Write(info)
	if err != nil {
		return Errors.New(err)
	}
	return nil
}
