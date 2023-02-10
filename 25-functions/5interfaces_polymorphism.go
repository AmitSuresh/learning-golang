package main

import "fmt"

type person struct {
	first, last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() { // receiver will attach func speak() with type. so that any value of that type has access to the func.
	fmt.Println("I am", s.first, s.last, "secret agent is called")
}

func (p person) speak() { // 
	fmt.Println("I am", p.first, p.last, "person is called")
}

// func (r receiver) identifier(parameters) (return(s)) { code... }

type human interface { // keyword identifier type
	speak()	// whatever type is used in this method, is now the type of this interface. so the original value is of its original type as well as this interface type.
}
// type identifier interface {} // represents everything as they can all implement empty interface. every value implemets empty interface.

func bar(h human) {
	switch h.(type){	//use of .(type) outside type switch IS NOT POSSIBLE.
	case person:
		fmt.Println("\nType is person.", h.(person).first)
	case secretAgent:
		fmt.Println("\nType is secretAgent.", h.(secretAgent).first)
	default: fmt.Printf("\nThe type is:%T\n", h)
	}
	fmt.Printf("Human is called:%v, Type:%T\n", h, h)
}

type hotdog int

func main() {

	p1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	fmt.Println(p1)

	p1.speak()

	p2 := secretAgent{
		person: person{
			"Holy",
			"Cow",
		},
		ltk: false,
	}

	p2.speak()

	p3 := person{
		"sad",
		"no",
	}
	fmt.Println(p3)

	bar(p1)
	bar(p2)
	bar(p3)

	// conversion

	var x hotdog = 42
	fmt.Printf("\nValue: %v, Type: %T\n", x, x)

	var y int
	y = int(x)
	fmt.Printf("\nValue: %v, Type: %T\n", y, y)

}