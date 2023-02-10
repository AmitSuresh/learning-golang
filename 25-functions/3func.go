package main

import "fmt"

func main() {
	defer foo()	//defered when main exits.
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
