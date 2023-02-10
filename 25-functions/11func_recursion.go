package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(looping_factorial(4))
	fmt.Println(loopFact(4))
}

func factorial(n int) int {
	switch {
	case n == 1:
		fmt.Println(n)
		return 1
	default:
		fmt.Println(n)
		return n * factorial(n-1)
	}
}

func looping_factorial(n int) int {
	fmt.Println("outside before loop:n:",n)
	for i:=(n-1);i>0;i-- {
		fmt.Println("before n:", n)
		fmt.Println("i:", i)
		n *= i
		fmt.Println("after n:", n)
	}
	fmt.Println("outside after loop:n:",n)
	return n
}

func loopFact(n int) int {	// course code.
	total := 1
	for ; n > 0; n-- {		// init left empty as it is not needed.
		total *= n
	}
	return total
}