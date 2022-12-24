package main

import (
	"fmt"
)

const (
	_ = iota
	//kilob = 1024
	kilob = 1 << (iota * 10)
	//megab = 1024 * kilob
	megab = 1 << (iota * 10)
	//gigab = 1024 * megab
	gigab = 1 << (iota * 10)
)

func main() {
	fmt.Printf("\n\t%d\t\t%b\n\t%d\t\t%b\n\t%d\t%b", kilob, kilob, megab, megab, gigab, gigab)
	x := 4
	fmt.Printf("\n\t%d\t\t%b", x, x)

	y := x << 1
	fmt.Printf("\n\t%d\t\t%b", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	
	fmt.Printf("\n\t%d\t\t%b\n\t%d\t\t%b\n\t%d\t%b", kb, kb, mb, mb, gb, gb)
}