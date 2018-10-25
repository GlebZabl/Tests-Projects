package main

import (
	. "Tests-Projects/bifit/structs"
	"fmt"
	"strconv"
)

func main() {
	parser:= FileParser{FieldSize:2}
	err,result := parser.Parse("./source.txt")

	if err == nil{
		printOut(result)
	}
}


func printOut(data [][]int) {
	for hIndex := 0; hIndex < len(data); hIndex++ {
		for vIndex := 0; vIndex < len(data); vIndex++ {
			var temp string
			if data[hIndex][vIndex] == -1{
				temp = "X"
			}else {
				temp = strconv.Itoa(data[hIndex][vIndex])
			}
			fmt.Print("  " + temp + "  ")
		}
		fmt.Println()
	}
}

