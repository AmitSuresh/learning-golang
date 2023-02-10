package main

import "fmt"

func main() {
	s1 := foo()

	fmt.Println(s1)
	s2 := func() int {
		return 1
	}()
	fmt.Printf("\n%v, %T\n\n", s2, s2)

	s3 := bar(5)
	fmt.Println(s3())

	fmt.Println(bar(5)())
}

func foo() string {
	s := "Hello world."
	return s
}

func bar(s int) func() int {	//func bar(s int) (func() int) {
	return func() int {
		return s + 10
	}
}
