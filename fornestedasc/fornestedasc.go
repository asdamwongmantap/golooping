package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {

		for k := 1; k <= i; k++ {
			fmt.Print("* ")

		}
		fmt.Println("")
	}
}
