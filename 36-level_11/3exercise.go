package main

import (
	"fmt"
)

type customErr struct {
	w1 string
}

func (c customErr) Error() string {
	return fmt.Sprintf("\nError : %v\n", c.w1)
}

func foo(e error) {
	fmt.Println("foo\n", e.(customErr).w1, "\n", e) // Assertion : e.(customErr).w1
}

func main() {
	some := customErr{
		"ss",
	}
	foo(some)
}
