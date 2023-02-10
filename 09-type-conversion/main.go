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
	fmt.Printf("value of a:\"%v\", value of b:\"%v\", value of c:\"%v\"\n\n",a,b,c)
	//a = c This assignment is not possible as they both belong to 2 different types.
	fmt.Printf("value of c:\"%v\",type of c:\"%T\", value of b:\"%v\",type of b:\"%T\"\n\n",c,c,b,b)
	a = int(c)
	d = string(b)
	fmt.Printf("value of a:\"%v\",type of a:\"%T\", value of d:\"%v\",type of d:\"%T\"\n\n",a,a,d,d)
}