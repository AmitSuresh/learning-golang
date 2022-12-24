package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Here.\n")
	modulo := 43 % 40
	fmt.Print(modulo)

	x := 1
	for {
		x++
		if x > 100{
			break
		}
		if x %2 != 0 {
			continue
		}
		fmt.Printf("Even numbers only: %v", x)
	}
}