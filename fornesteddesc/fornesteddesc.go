package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		for k := i; k < 5; k++ {
			fmt.Print("* ")

		}
		fmt.Println("")
	}
}
