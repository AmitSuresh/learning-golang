package main

import "fmt"

func main() {
	var x [5]int

	fmt.Printf("The size of x is: %v\n", len(x))

	fmt.Printf("the value of x: %v\n", x)

	count := 0

	for i := range x {
		count++
		x[i] = count
	}

	fmt.Printf("the value of x after loop: %v", x)
}
