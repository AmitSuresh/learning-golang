package main

import "fmt"

func main() {
	x := 0
	fmt.Printf("Years alive:")
	for i := 1997; i <= 2022; i++ {
		x++
		fmt.Printf("\n\t%v", i)
	}
	fmt.Printf("\nAge: %v", x)
}