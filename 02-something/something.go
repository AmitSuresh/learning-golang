package main

import "fmt"

func main() {
	fmt.Println("Hello.")
	foo()
	fmt.Println("after foo")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	n, e := fmt.Println("I'm in foo")
	fmt.Println(n)
	fmt.Println(e)
}

func bar() {
	n, _ := fmt.Println("and then we exit")
	fmt.Println(n)
}
