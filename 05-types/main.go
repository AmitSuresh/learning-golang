package main

import "fmt"

var y = 14
var z = "Superman, batsie"
var a = `Let us check out "Raw String Literal"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}