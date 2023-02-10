package main

import (
	"fmt"
)

func main() {
	returned_func_variable := returning_func("Calling from variable.")
	numbers_slice := []int{10, 20, 30, 40, 50}
	sum := returned_func_variable(numbers_slice...)
	fmt.Printf("Sum of numbers_slice%v: %v", numbers_slice, sum)
}

func returning_func(st string) func(num ...int) int {
	fmt.Println(st)
	return func(num ...int) int {
		fmt.Println("returned this func.")
		sum_of_num := 0
		for _, v := range num {
			sum_of_num += v
		}
		return sum_of_num
	}
}
