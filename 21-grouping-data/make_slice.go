package main

import "fmt"

func main() {
	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[1] = 253 // replaces 0 in 1 index with 253.
	fmt.Println(x)
	//x[10] = 21 // not possible. index outside the range.
	x = append(x, 21) // increases capacity of array to 20 and adds value 21 to the index 10. doubles.
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	o := 0
	for i := range x {
		o++
		x[i] = o
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21) // increases capacity of array to 40. doubles.
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
