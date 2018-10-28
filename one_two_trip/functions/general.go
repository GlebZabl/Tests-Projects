package functions

import (
	"math/rand"
	"strconv"
)

//c 5ти процентной вероятностью возвращает ошибку
func GetError() bool {
	return false
	random := rand.Intn(19)
	return random == 0
}

//генерирует рандомную строчку
func GetRandomString() string {
	return strconv.Itoa(rand.Int()) + strconv.Itoa(rand.Int()) + strconv.Itoa(rand.Int())
}

//читаем очередь ошибок
func ReadErrors() {

}

//возвращает режим работы
func GetMode() bool {
	return false
}
