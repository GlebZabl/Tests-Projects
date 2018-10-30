package main

import "fmt"

func main() {

	a := [][]int{[]int{17, 16, 15, 14, 13}, []int{18, 5, 4, 3, 12}, []int{19, 6, 1, 2, 11}, []int{20, 7, 8, 9, 10}, []int{21, 22, 23, 24, 25}}

	readLeftDown(a, readRightUp(a, readLeftDown(a, readRightUp(a, [2]int{2, 2}, 1),1), 3),3)
}

func readMatrix(matrix [][]int, courIndexes [2]int, age int) {
	courIndexes = readRightUp(matrix, courIndexes, age)
	age++
	courIndexes = readLeftDown(matrix, courIndexes, age)
	age++
	readMatrix(matrix, courIndexes, age)
}

func readRightUp(matrix [][]int, courIndexes [2]int, age int) [2]int {

	for i := courIndexes[1]; i <= courIndexes[1]+age; i++ {
		fmt.Print(" ")
		fmt.Print(matrix[courIndexes[0]][i])
	}
	for i := courIndexes[0] - 1; i >= courIndexes[0]-age; i-- {
		fmt.Print(" ")
		fmt.Print(matrix[i][courIndexes[1]+age])
	}

	courIndexes = [2]int{courIndexes[0] - age, courIndexes[1] + (age - 1)}
	return courIndexes
}

func readLeftDown(matrix [][]int, courIndexes [2]int, age int) [2]int {

	for i := courIndexes[1]; i >= courIndexes[1]-age; i-- {
		fmt.Print(" ")
		fmt.Print(matrix[courIndexes[0]][i])
	}
	for i := courIndexes[0]+1;i<=courIndexes[0]+age;i++{
		fmt.Print(" ")
		fmt.Print(matrix[i][courIndexes[1]-age])
	}
	courIndexes = [2]int{courIndexes[0] + (age+1), courIndexes[1] - age}
	return courIndexes
}
