package main

import (
	"fmt"
)

func main() {
	a, b := enclosed_variable(), enclosed_variable()
	fmt.Println("a:\t", a(), "\tb:\t", b())
	fmt.Println("a:\t", a())
	fmt.Println("a:\t", a(), "\tb:\t", b())
	fmt.Println("a:\t", a())
	fmt.Println("a:\t", a(), "\tb:\t", b())
	fmt.Println("a:\t", a(), "\tb:\t", b())
}

func enclosed_variable() func() int {
	var enclosed_var int = 1
	return func() int {
		enclosed_var *= 5
		return enclosed_var
	}
}
