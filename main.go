package main

import (
	"fmt"
	"main/operations"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage error: please provide exactly 1 file name as argument.")
		return
	}

	file := os.Args[1]
	input, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	tetrominoes := operations.Split(input)
	fmt.Println(tetrominoes)
	if operations.Checker(tetrominoes) {
		fmt.Println("good")
	} else {
		fmt.Println("Error")
	}

}
