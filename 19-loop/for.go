// `for` is Go's only looping construct. Here are
// some basic types of `for` loops.

package main

import "fmt"

func main() {

	fmt.Printf("1.\n\n")
	// The most basic type, with a single condition.
	
	for i := 1; i <= 3; {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Printf("\n\n2.\n\n")
	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Printf("\n\n3.\n")
	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}
	fmt.Printf("\n\n4.\n\n")
	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	
	x := 15
	for x<11 {
		fmt.Println(x)
		//break
	}
	y := 15
	for {
		if y>25 {	
			fmt.Println(y)
			break
		}
		y++
		fmt.Printf("incremented y: %v\n", y)
	}
}
