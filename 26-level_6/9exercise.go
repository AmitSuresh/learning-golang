package main

import (
	"fmt"
)

func main() {
	str, returned_func_variable := returning_func("Calling from variable.")
	fmt.Println("Printing returned string str= ", str)

	numbers_slice := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55}

	sum_of_even_numbers_of_slice, even_numbers_of_slice_before_addition := returned_func_variable(sum_of_slice, numbers_slice...)

	fmt.Printf("\nSum of Even Numbers of numbers_slice%v:\t%v\nEven Numbers before addition:\t%v", numbers_slice, sum_of_even_numbers_of_slice, even_numbers_of_slice_before_addition)
}

func returning_func(st string) (string, func(callback_sum_func func(num ...int) int, num ...int) (int, []int)) { //RETURN STATEMENT: (func(callback_sum_func func(num ...int) int, num ...int) []int )
	fmt.Println(st)
	return st, func(callback_sum_func func(num ...int) int, num ...int) (int, []int) {
		even_numbers := make([]int, 0, 0)
		for _, v := range num {
			if v%2 == 0 {
				even_numbers = append(even_numbers, v)
			}
		}
		return callback_sum_func(even_numbers...), even_numbers // this is now an integer. as (sum_of_slice(num ...int) int) returns int
	}
}

func sum_of_slice(num ...int) int {
	total := 0
	for _, v := range num {
		total += v
	}
	return total
}
