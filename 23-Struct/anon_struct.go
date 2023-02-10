package main

import "fmt"

type foo int

var y foo

const x = 33

func main() {
	p1 := struct { // anonymous struct.
		first, last string
		age         int
	}{
		first: "James",
		last:  "Bond",
		age:   54,
	}

	fmt.Println(p1)

	y = x

	fmt.Printf("Type %T, %T, %T", int(y), y, x)
}
