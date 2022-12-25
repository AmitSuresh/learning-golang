package main

import "fmt"

func main() {
	x := []int{25, 342, 52, 56, 7, 34, 64, 43, 2, 76, 76, 6}
	h := append(x[1:3], x[6:11]...) // first argument takes a slice, the second takes values.
	//use []Type{value1,value2,value3} in first. value1,value2,value3 in second.
	// if x and y are slices. first argument = x, second argument = y...
	// when assigning new value to slice on a new line, make sure to use the same slice in first argument. this is not python.
	fmt.Printf("\nh is: %v\n", h)
}
