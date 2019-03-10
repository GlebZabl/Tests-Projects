package logger

import (
	"test_project/envirement/errors"
)

type fullPrinter struct {
	external *mainLogger
	file     *filePrinter
	console  *consolePrinter
}

func (f *fullPrinter) setExternal(logger *mainLogger) {
	f.external = logger
}

func (f *fullPrinter) print(info string) error {
	_ = f.console.print(info)
	if err := f.file.print(info); err != nil {
		return Errors.New(err)
	}
	return nil
}

func newFullPrinter() printer {
	cnslPrinter := newConsolePrinter()
	flPrinter := newFilePrinter()
	return &fullPrinter{
		file:    flPrinter.(*filePrinter),
		console: cnslPrinter.(*consolePrinter),
	}
}
