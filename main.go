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

		mode = !mode

		for j := 0; j < number; j++ {

			if mode {
				fmt.Printf("# ")
			} else {
				fmt.Printf(" #")
			}
		}
		fmt.Printf("\n")
	}

}
