package main

import "fmt"

func main() {
	foo()
	func() {	// func(parameters){code}()		parenthesis with arguments at the end
		fmt.Println("anon func. or anon self executing func.")
	}()
	func(a int) {
		fmt.Println("anonymous func with parameter. We are passing in 1 argument.",a)
	}(2)
}

func foo() {
	fmt.Println("foo")
}
