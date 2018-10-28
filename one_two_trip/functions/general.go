package functions

import (
	"math/rand"
	"strconv"
)

//c 5ти процентной вероятностью возвращает ошибку
func GetError() bool {
	random := rand.Intn(19)
	return random == 0
}

func GetRandomString() string  {
	return strconv.Itoa(rand.Int())+strconv.Itoa(rand.Int())+strconv.Itoa(rand.Int())

}

func ReadErrors()  {

}

func GetMode() bool  {
	return false
}