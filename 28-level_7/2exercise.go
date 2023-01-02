package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	james := person{
		name: "bond",
	}
	fmt.Printf("person: %v\n", james)
	changeMe(&james, "James Bond")
	fmt.Printf("changed person: %v\n", james)

	v := person{
		"valentino",
	}
	fmt.Printf("valentino: %v\n", v)
	changeMeWithoutParam(v, "James Bond")
	fmt.Printf("changed person: %v\n", v)

}

func changeMe(p *person, str string) {
	(*p).name = str
	fmt.Println()
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println()
}

func changeMeWithoutParam(p person, str string) {
	(p).name = str
	fmt.Println()
	fmt.Println(&p)
	fmt.Println(p)
	fmt.Println()
}