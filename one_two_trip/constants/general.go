package constants

import (
	"bufio"
	"os"
	"strconv"
)

const confPath = "./config.txt"

var (
	RedisConString string
	RedisPort      string
	RedisDbNumber  int64
	TasksQueueName string
	ErrQueueName   string
	ChanelName     string
	VoteChanel     string
	VoteNotifier   string
)

//устанавливаем константы
func init() {
	RedisConString, RedisPort,RedisDbNumber, TasksQueueName, ErrQueueName, ChanelName, VoteChanel, VoteNotifier = loadFromConfig(confPath)
}

//чтение из файла
func loadFromConfig(path string) (string, string, int64, string, string, string, string, string) {
	result := *new([]string)
	file, err := os.Open(path)
	if err != nil {
		return "", "", 0, "", "", "", "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	dbnmbr, err := strconv.Atoi(result[2])
	if err != nil {
		return "", "", 0, "", "", "", "", ""
	}
	return result[0], result[1], int64(dbnmbr), result[3], result[4], result[5], result[6], result[7]
}
