package main

import (
	"Tests-Projects/bifit/structs"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const filePath = "./source.txt"

func main() {

	err, size := getSize(filePath)
	if err{
		println("cant load source file, sorry")
		return
	}

	parser := structs.FileParser{FieldSize: size}
	err, result := parser.Parse(filePath)
	if !err{
		printOut(result)
	}else{
		println("wrong field format, checkout what it must looks like, from readme ")
	}
}

//функция для корректного вывода матрицы в консоль
func printOut(data [][]int) {
	for hIndex := 0; hIndex < len(data); hIndex++ {
		for vIndex := 0; vIndex < len(data); vIndex++ {
			var temp string
			if data[hIndex][vIndex] == -1 {
				temp = "X"
			} else {
				temp = strconv.Itoa(data[hIndex][vIndex])
			}
			fmt.Print("  " + temp + "  ")
		}
		fmt.Println()
	}
}

//смотрим размер одной стороны поля
func getSize(path string) (bool,int) {
	result := -1
	file, err := os.Open(path)
	if err != nil {
		return true, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line:=scanner.Text()
	for i := range line{
		if line[i:i+1] == "O" || line[i:i+1] == "X"{
			result++
		}
	}
	return false,result
}
