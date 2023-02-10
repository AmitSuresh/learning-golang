package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("The values of x, y, z: %v, %v, %v\n", x, y, z)
	fmt.Printf("The types of x, y, z: %T, %T, %T", x, y, z)
}
