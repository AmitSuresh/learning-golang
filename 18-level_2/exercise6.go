package main

import (
	"fmt"
)

const (
	a = 2022 + iota
	b 
	c 
	d
)

func main() {
	fmt.Printf("Year: %v, Year %v, Year %v, Year %v", a, b, c, d)
}