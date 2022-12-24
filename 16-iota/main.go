package main

import "fmt"

const (
	a = iota
	b 
	c 
)

const (
	d = iota
	e 
	f 
)
func main() {
	fmt.Printf("\n\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v", a, a, b, b, c, c, d, d, e, e, f, f)
}