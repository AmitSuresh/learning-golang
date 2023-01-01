package main

import (
	"fmt"
)

func main() {
	slice_int := []int{1, 2, 3, 4}
	fmt.Println(foo(slice_int...))
	fmt.Println(bar(slice_int))
}

func foo(num ...int) int {
	sum_of_num := 0
	for _, v := range num {
		sum_of_num += v
	}
	return sum_of_num
}

func bar(slice_of_int []int) int {
	defer ex3()		// exercise 3
	sum_of_num := 0
	for _, v := range slice_of_int {
		sum_of_num += v
	}
	return sum_of_num
}

func ex3() {
	fmt.Println("defer is run.")
}
