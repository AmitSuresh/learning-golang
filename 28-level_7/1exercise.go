package main

import (
	"fmt"
)

func main() {
	a := 12
	fmt.Printf("address of %v is %v", a, &a)
}
