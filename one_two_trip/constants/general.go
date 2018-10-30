package constants

import (
	"bufio"
	"os"
)

const confPath = "./config.txt"

var (
	RedisConString string
	TasksQueueName string
	ErrQueueName   string
	ChanelName     string
	VoteChanel     string
	VoteNotifier   string
)

//устанавливаем константы
func init() {
	RedisConString, TasksQueueName, ErrQueueName, ChanelName, VoteChanel, VoteNotifier = loadFromConfig(confPath)
}

//чтение из файла
func loadFromConfig(path string) (string, string, string, string, string, string) {
	result := *new([]string)
	file, err := os.Open(path)
	if err != nil {
		return "", "", "", "", "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result[0], result[1], result[2], result[3], result[4], result[5]
}
