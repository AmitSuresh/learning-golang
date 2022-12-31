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
	fmt.Println("I am", s.first, s.last)
}

// func (r receiver) identifier(parameters) (return(s)) { code... }

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

}
