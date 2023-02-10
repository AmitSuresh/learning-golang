package main

import "fmt"

const (
	a = iota	//0
	b 
	c = 0	//skips
	d = iota	//3
	e
	f 
)

const (
	g = iota
	h 
	i 
)
func main() {
	fmt.Printf("%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v\t%T\t%v", a, a, b, b, c, c, d, d, e, e, f, f, g, g, h, h, i, i)
}