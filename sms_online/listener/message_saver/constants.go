package message_saver

import (
	"bufio"
	"os"
)

const (
	confPath = "./config.txt"
)

var (
	DBConnectionString    string //строка для подключения к бд
	QueueConnectionString string //строка для подключения к менеджеру очередей
	QueueName             string //название очереди
)

func init() {
	DBConnectionString, QueueConnectionString, QueueName = loadConfig(confPath)
}

//грузим переменные из файла конфига
func loadConfig(path string) (string, string, string) {
	result := *new([]string)
	file, err := os.Open(path)
	if err != nil {
		return "", "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result[0], result[1], result[2]
}
