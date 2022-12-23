package main

import "fmt"

type customtype string
var x customtype
var y string 

func main() {
	fmt.Printf("Value and Type of x are:%v and %T", x, x)
	x = "Hello, World."
	fmt.Println(x)
	y = string(x)
	fmt.Printf("Value and Type of y are:%v and %T", y, y)
}