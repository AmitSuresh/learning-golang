package main

import "fmt"

func main() {
	if x := 17; x == 2 {
		fmt.Println("001")
	}
	fmt.Println("statement 1"); fmt.Println("statement 2")//; fmt.Println(x)	not possible to call x outside the if loop.
	y := 17
	if y == 2 {
		fmt.Println("001")
	}
	fmt.Println(y)
	fmt.Printf("\n\nnext\n\n")
	x := 23
	if x == 3 {
		fmt.Printf("value is %v", x)
	} else if x == 22 {
		fmt.Printf("value is %v", x)
	} else {
		fmt.Printf("value is %v", x)
	}
	fmt.Printf("\n\nnext\n\n")
	for i := 0; i < 100; i ++ {
		if i % 7 == 0 {
			fmt.Println(i)
		}
	}
}