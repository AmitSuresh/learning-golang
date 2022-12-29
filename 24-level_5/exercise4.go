package main

import "fmt"

func main() {
	x := struct {
		who    string
		angry  bool
		reason string
	}{
		"Amit", true, "Taking too long",
	}

	fmt.Printf("Who is angry?\n\t%v\nIs he angry now?\n\t%v\nWhy is he angry?\n\t%v", x.who, x.angry, x.reason)
}
