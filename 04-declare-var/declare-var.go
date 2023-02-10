package main

import "fmt"

// z := 60 // it is not possible to use := outside function.
var z = 60	// scope of z is package level
var a int
func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)
	var y = 43
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	foo()
}

func foo() {
	fmt.Println("again:",z)
}