package main

import "fmt"

func main() {
	foo() //when calling func, they are arguments.
	bar("Johnny")
	s1 := woo("Geralt")
	fmt.Println(s1)
	st, bl := uhoh("Ian", "Fleming")
	fmt.Println(st)
	fmt.Println(bl)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }	// when defining func, they are parameters.

func foo() {
	fmt.Println("Hello from foo.")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func uhoh(first, last string) (string, bool) {
	str := fmt.Sprint(first," ", last, `, says "Hello"`)
	bl := true
	return str, bl
}
