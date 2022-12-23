package main

import "fmt"

var a int
type johnny string
type integertype int
var d string

func main() {
	a = 7
	var b johnny 
	b = "Eat healthy"
	var c integertype
	c = 14
	fmt.Printf("value of a:\"%v\", value of b:\"%v\", value of c:\"%v\"",a,b,c)
	//a = c This assignment is not possible as they both belong to 2 different types.
}