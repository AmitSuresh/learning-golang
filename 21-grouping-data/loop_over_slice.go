package main

import "fmt"

func main() {
	x := []int{25, 342, 52, 56, 7}
	fmt.Println(len(x))
	fmt.Printf("%v", x[3])
	fmt.Println()
	fmt.Println(cap(x))
	fmt.Println()
	fmt.Println()

	for i, v := range x { // i and v is commonly used.
		fmt.Printf("index i: %v, value in index: %v\n", i, v)
	}
}
