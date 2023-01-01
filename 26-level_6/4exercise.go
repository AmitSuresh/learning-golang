package main

import (
	"fmt"
)

type person struct {
	first, last string
	age         int
}

func main() {
	johnny := person{
		"Johnny",
		"Silverhand",
		45,
	}
	johnny.speak()
}

func (prsn person) speak() {
	fmt.Println("\nName of the person:\t", prsn.first, "\nAge:\t\t\t", prsn.age)
}
