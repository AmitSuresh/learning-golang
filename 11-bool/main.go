package main

import (
	"fmt"
)

func main() {
	a := 8
	b := 10
	fmt.Println(a == b)
	if a<b {
		fmt.Println(a<b)
	} else if a>b {
		fmt.Println(a>b)
	}
}