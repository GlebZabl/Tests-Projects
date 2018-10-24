package message_saver

import (
	"fmt"
	"io/ioutil"
)

const(
	confPath = "./config.txt"
)

var(
	DBConnectionString string
)

func init() {
	DBConnectionString = loadDbConectionString(confPath)
}

func loadDbConectionString(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
