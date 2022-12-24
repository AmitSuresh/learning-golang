package main

import "fmt"

func main() {
	switch "Holy" {
	case "Unholy":
		fmt.Println("001")
	case "Insanity":
		fmt.Println("002")
	case "Lovely", "Holy", "Crazy":
		fmt.Println("before holy... COW")
	case "holy":
		fmt.Println("Cow")
	case "Hello":
		fmt.Println("005")
	default:	// if all the above cases are false, this will fire.
		fmt.Println("This is default.")
	}
}
