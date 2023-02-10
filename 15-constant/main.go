package main

import "fmt"

const (
	a int = 42
	b string = "Jones"
	c float32 = 42.23
)

func main() {
	fmt.Printf("type,value:\n\t%T\t%v\n\t%T\t%v\n\t%T\t%v", a, a, b, b, c, c)
}