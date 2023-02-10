package main

import "fmt"

func main() {
	fmt.Println(sum([]int{2, 5, 23, 4, 65, 56}...))
	fmt.Println(even(sum, []int{2, 5, 23, 4, 65, 56}...))
	fmt.Println(odd(sum, []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,}...))
}

func sum(num1 ...int) int {
	total := 0
	for _, v := range num1 {
		total += v
	}
	return total
}

func even(f func(num1 ...int) int, num2 ...int) int {
	even_num := make([]int, 0, 10)
	for _, v := range num2 {
		if v%2 == 0 {
			even_num = append(even_num, v)
		}
	}
	return f(even_num...)
}

func odd(f func(num1 ...int) int, num3 ...int) int {
	odd_num := make([]int, 0, 20)
	for _, v := range num3 {
		if v%2 == 1 {
			odd_num = append(odd_num, v)
		}
	}
	fmt.Println(odd_num)
	return f(odd_num...)
}
