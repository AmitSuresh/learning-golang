package main

import "fmt"

func main() {
	x := sum("some", 1, 2, 3, 4, 5, 6)
	fmt.Println("Sum of the integers:", x) //() is needed even if there's no arguments.
	fmt.Println()
	a := []int{42, 63, 34, 67, 4, 2}
	y := sum("some", a...) //unpack the slice and pass it.
	fmt.Println("Sum of y:", y)
	fmt.Println()

	z := sum("some") //variadic parameters can take nothing.
	fmt.Println("Sum of z:", z)
	fmt.Println()

	b := sum("some", 3) //variadic parameters can take nothing.
	fmt.Println("Sum of b:", b)
	fmt.Println()
}

func sum(s string, x ...int) int { //x stores infinite number of int parameters as a slice of int.
	//func sum(x ...int, s string) int { //can only use ... with final parameter in list
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Printf("%T\n", x)

	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

// func (r receiver) identifier(parameter(s)) (return(s)){ ... }
