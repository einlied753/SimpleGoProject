package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Please, enter the board dimension number...")
	fmt.Scan(&number)

	mode := false

	for i := 0; i < number; i++ {

		for j := 0; j < number; j++ {

			mode = !mode

			if mode {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}
