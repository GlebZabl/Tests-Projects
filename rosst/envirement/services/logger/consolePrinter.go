package logger

import (
	"Tests-Projects/rosst/envirement/errors"
	"fmt"
)

type consolePrinter struct {
	external *mainLogger
}

func (c *consolePrinter) setExternal(logger *mainLogger) {
	c.external = logger
}

func (c *consolePrinter) print(source string) error {
	_, err := fmt.Println(source)
	return Errors.New(err)
}

func newConsolePrinter() printer {
	return new(consolePrinter)
}
