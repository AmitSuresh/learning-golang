package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func (pers *person) speak() {
	fmt.Println("", pers.Name, "- ", pers.Age, " says Hello")
}

type human interface {
	speak()
}

func saySomething(hum human) {
	hum.speak()
}

func main() {

	Johnny := person{
		"Johnny Silverhand",
		44,
	}

	//saySomething(Johnny)	//This won't work as Johnny is not a pointer
	saySomething(&Johnny)
}
