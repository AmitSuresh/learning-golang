package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a(), b(), incrementor()())
	fmt.Println(a(), b(), incrementor()())
	fmt.Println(a(), b(), incrementor()())
	fmt.Println(a(), b(), incrementor()())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
