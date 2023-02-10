package main

import (
	"fmt"
)

func main() {
	a := 7
	fmt.Printf("%d, %b, %#x\n", a, a, a)
	b := 1 << a
	fmt.Printf("%d, %b, %#x", b, b, b)
}