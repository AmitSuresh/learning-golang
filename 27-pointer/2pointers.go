package main

import (
	"fmt"
)

func main() {
	a := foo(77)
	fmt.Println()
	fmt.Println("*a", *a)
	fmt.Println("a", a)
	fmt.Println()
	b := bar(a, 190)
	fmt.Println()
	fmt.Println("*b", *b)
	fmt.Println("b", b)
	fmt.Println()
	c := foo(*b)
	fmt.Println()
	fmt.Println("*c", *c)
	fmt.Println("c", c)
	fmt.Println()
	d := bar(c, 3)
	fmt.Println()
	fmt.Println("*d", *d)
	fmt.Println("d", d)
	fmt.Println()
	e := bar(d, 23)
	fmt.Println()
	fmt.Println("*e", *e)
	fmt.Println("e", e)
	fmt.Println()
	f := bar(e, 1)
	fmt.Println()
	fmt.Println("*f", *f)
	fmt.Println("f", f)
	fmt.Println()
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("f", f)
	fmt.Println()
	fmt.Println("&a", &a)
	fmt.Println("&b", &b)
	fmt.Println("&c", &c)
	fmt.Println("&d", &d)
	fmt.Println("&e", &e)
	fmt.Println("&f", &f)
	fmt.Println()
	fmt.Println("*a", *a)
	fmt.Println("*b", *b)
	fmt.Println("*c", *c)
	fmt.Println("*d", *d)
	fmt.Println("*e", *e)
	fmt.Println("*f", *f)

}

func foo(x int) *int {
	return &x
}

func bar(y *int, z int) *int {
	fmt.Printf("\ny is:%v\nz is:%v\n", *y, z)
	*y = z
	return y
}
