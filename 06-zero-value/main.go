package main

import (
	"fmt"
)

var a string
var b int
var c bool

func main() {
	fmt.Printf("%T\n",a)
	fmt.Println("value of a=",a)
	fmt.Printf("%T\n",b)
	fmt.Println("value of b=",b)
	fmt.Printf("%T\n",c)
	fmt.Println("value of c=",c)
	a = "Johnny Silverhand"
	fmt.Printf("\n%T\n",a)
	fmt.Println("value of a=",a)
	b = 7
	fmt.Printf("%T\n",b)
	fmt.Println("value of b=",b)
	c = true
	fmt.Printf("%T\n",c)
	fmt.Println("value of c=",c)
	fmt.Printf("value of c=%t",c)
}