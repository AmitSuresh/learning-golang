package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		if x % 4 == 0 {
			fmt.Printf("Remainder of %v/4\t:\t%v\n", x, x%4)
		}
	}
}