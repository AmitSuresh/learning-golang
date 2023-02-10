package main

import "fmt"

func main() {
	switch {	// missing switch expression = true
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

	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}