package main

import "fmt"

func main() {
	if x := 13; x < 0 {
		fmt.Printf("%v is less than 0.", x)
	}else if x == 0 {
		fmt.Printf("%v is equal to 0.", x)
	}else {
		fmt.Printf("%v is greater than 0.", x)
	}
}
