package main

import "fmt"

func main() {
	switch {
	case 3 > 5:
		fmt.Printf("3 is greater than 5.")
	case 3 < 0:
		fmt.Printf("3 is less than 0.")
	case 3 > 0:
		fmt.Printf("3 is greater than 0.")
	case 3 == 0:
		fmt.Printf("3 is equal to 3.")
	default:
		fmt.Printf("default")
	}
}
