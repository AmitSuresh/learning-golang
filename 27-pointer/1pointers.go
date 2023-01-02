package main

import (
	"fmt"
)

func main() {
	a := 23
	fmt.Printf("value:%v, type:%T, address: %v, type of address: %T\n", a, a, &a, &a)

	b := &a
	fmt.Printf("\nval:%v, type:%T,address:%v, type:%T\n", b, b, &b, &b)
	fmt.Printf("What is stored in b?: %v", *b)

	var c **int
	c = &b
	fmt.Printf("\nval:%v, type:%T,address:%v, type:%T\n", c, c, &c, &c)
	fmt.Printf("\nWhat is stored in c?: %v\n", *c)

	d := 345
	fmt.Printf("\nvalue of d: %v, address: %v, value in address: %v\n", d, &d, *&d)
	// &x = address. *&x = value stored in address (&x)
	*b = 234
	fmt.Printf("value:%v, type:%T, address: %v, type of address: %T\n", a, a, &a, &a)
}
