package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
	z := "Hello there."
	if x < y {
		fmt.Println("x=",x,z)
	}
}

/* func compare() {
	if x < y {
		fmt.Println("x=%q",x)
	}
} */