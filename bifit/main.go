package main

import (
	. "Tests-Projects/bifit/structs"
	"fmt"
)

func main() {
	err, result := loadFile("source.txt")

	if err != nil || !check(result) {
		fmt.Println("wrong field format!")
		return
	}

	counter := Counter{Source: result}
	result = counter.Count()
	printOut(result)
}

func loadFile(path string) (error, [][]string) {
	parser := new(FileParcer)
	err, result := parser.Parse(path)
	return err, result
}

func printOut(data [][]string) {
	for hIndex := 0; hIndex < len(data); hIndex++ {
		for vIndex := 0; vIndex < len(data); vIndex++ {
			fmt.Print(" " + data[hIndex][vIndex] + "")
		}
		fmt.Println()
	}
}

func check(data [][]string) bool {
	for index := range data {
		if len(data[index]) != len(data) {
			return false
		}
	}
	return true
}
