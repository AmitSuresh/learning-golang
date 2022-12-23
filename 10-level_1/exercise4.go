package main

import "fmt"

type underlyingtype int
var x underlyingtype
type T1 string
type T2 T1
type T3 []T2
type T4 T3

func main() {
	fmt.Printf("Value of x=%v, Type of x=%T\n",x ,x)
	x = 42
	fmt.Printf("Types:\t%T\t%v", x, x)
	var a T1 = "as"
	var b T2 = "as"
	var c T3 = []T2{"A", "B", "C"}
	var d T4 = []T2{"A", "B", "C"}
	fmt.Println(a)
	fmt.Println(x)
	fmt.Printf("Types:\t%T\t%T\t%T\t%T", a, b, c, d)
}