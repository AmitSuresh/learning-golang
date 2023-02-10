package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5,6,7,8,9,0}

	fmt.Printf("The size of x is: %v\n", len(x))

	fmt.Printf("the value of x: %v\n", x)


	for i,v := range x {
		fmt.Printf("Index: %v, Value: %v, length: %v, capacity: %v\n", i, v, len(x), cap(x))
	}

	fmt.Printf("the value of x after loop: %v, Type of slice: %T", x, x)
}
