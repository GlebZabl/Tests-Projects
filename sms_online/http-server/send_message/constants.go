package send_message

import (
	"bufio"
	"os"
)

const (
	confPath = "./config.txt"
)

var (
	QueueConnectionString string
	QueueName             string
	ErrMsg                string
)

func init() {
	QueueConnectionString, QueueName = loadQueueConectionString(confPath)
	ErrMsg = "service is temporary unavailable"
}

//загружаем из конфигурационного файла
func loadQueueConectionString(path string) (string, string) {
	result := *new([]string)
	file, err := os.Open(path)
	if err != nil {
		return "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result[0], result[1]
}
