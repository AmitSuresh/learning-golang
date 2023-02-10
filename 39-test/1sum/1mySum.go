package main

import "fmt"

func main() {
	fmt.Println("5+5=", mySum(5, 5))
}

func mySum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
