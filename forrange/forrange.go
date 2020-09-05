package main

import (
	"fmt"
)

func main() {

	var number = []int{1, 2, 3, 4, 5}

	for i, rows := range number {
		fmt.Println("Index ", i, " Is :", rows)
	}
}
