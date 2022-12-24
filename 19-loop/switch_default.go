package main

import "fmt"

func main() {
	switch {
	case 0 == 5:
		fmt.Println("001")
	case 1 == 5:
		fmt.Println("002")
	case 2 == 5:
		fmt.Println("003")
	case 3 == 5:
		fmt.Println("004")
	case 4 == 5:
		fmt.Println("005")
	default:	// if all the above cases are false, this will fire.
		fmt.Println("This is default.")
	}
}
