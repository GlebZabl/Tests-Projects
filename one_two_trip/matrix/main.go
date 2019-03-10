package main

import "fmt"

func main() {

	source := [][]int{
		{17, 16, 15, 14, 13},
		{18, 5, 4, 3, 12},
		{19, 6, 1, 2, 11},
		{20, 7, 8, 9, 10},
		{21, 22, 23, 24, 25},
	}
	/*
		source := [][]int{
			{5,  4,  3},
			{6,  1,  2},
			{7,  8,  9},
		}
		source := [][]int{
			{1},
		}*/

	result := new([]int)

	//находим центральные клетки
	startIndexes := [2]int{(len(source) - 1) / 2, (len(source) - 1) / 2}
	readMatrix(source, startIndexes, 1, result)
	fmt.Println(*result)
}

//рикурсивно читаем матрицу
func readMatrix(matrix [][]int, idx [2]int, age int, result *[]int) {
	if age > len(matrix) {
		return
	}
	readMatrix(matrix, readOneAge(matrix, idx, age, result), age+2, result)
}

//читаем 1 "груг" матрицы
func readOneAge(matrix [][]int, idx [2]int, age int, result *[]int) [2]int {
	return readLeftDown(matrix, readRightUp(matrix, idx, age, result), age, result)
}

//читаем вправо и вверх
func readRightUp(matrix [][]int, courIndexes [2]int, age int, result *[]int) [2]int {

	for i := courIndexes[1]; i <= courIndexes[1]+age; i++ {
		if i <= len(matrix)-1 {
			*result = append(*result, matrix[courIndexes[0]][i])
		} else {
			return [2]int{0, 0}
		}
	}
	for i := courIndexes[0] - 1; i >= courIndexes[0]-age; i-- {
		*result = append(*result, matrix[i][courIndexes[1]+age])
	}

	courIndexes = [2]int{courIndexes[0] - age, courIndexes[1] + (age - 1)}
	return courIndexes
}

//читаем влево и вниз
func readLeftDown(matrix [][]int, courIndexes [2]int, age int, result *[]int) [2]int {
	if age >= len(matrix)-1 {
		return [2]int{0, 0}
	}
	for i := courIndexes[1]; i >= courIndexes[1]-age; i-- {
		*result = append(*result, matrix[courIndexes[0]][i])
	}
	for i := courIndexes[0] + 1; i <= courIndexes[0]+age; i++ {
		*result = append(*result, matrix[i][courIndexes[1]-age])
	}
	courIndexes = [2]int{courIndexes[0] + (age + 1), courIndexes[1] - age}
	return courIndexes
}
