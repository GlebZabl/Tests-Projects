package main

import "fmt"

func main()  {

	a:=[][]int{[]int{1,2,3,4,5},[]int{6,7,8,9,10},[]int{11,12,13,14,15},[]int{16,17,18,19,20},[]int{21,22,23,24,25}}
	//idx:=readRightUp(a,[2]int{1,1},1)
	//idx = readLeftDown(a,idx,2)
	//idx = readRightUp(a,idx,3)
	//idx := readLeftDown(a,[2]int{0,2},3)
	//fmt.Print(a[idx[0]][idx[1]])
	fmt.Println(readRightUp(a,[2]int{3,1},3))
}

func readMatrix(matrix [][]int, courIndexes [2]int,age int)  {
	courIndexes = readRightUp(matrix, courIndexes,age)
	age++
	courIndexes = readLeftDown(matrix, courIndexes,age)
	age++
	readMatrix(matrix, courIndexes,age)
}

func readRightUp(matrix [][]int,courIndexes [2]int,age int) [2]int {

	for i:=courIndexes[1];i<=courIndexes[1]+age;i++{
		fmt.Print(" ")
		fmt.Print(matrix[courIndexes[0]][i])
	}
	for i:=courIndexes[0]-1;i>=courIndexes[0]-age;i--{
		fmt.Print(" ")
		fmt.Print(matrix[i][courIndexes[1]+age])
	}

	courIndexes = [2]int{courIndexes[0]-age,courIndexes[1]+(age-1)}
	return courIndexes
}

func readLeftDown(matrix [][]int,courIndexes [2]int,age int) [2]int {

	courIndexes = [2]int{courIndexes[0]+age,courIndexes[1]-(age-1)}
	return courIndexes
}

