package main

import (
	"fmt"
)

const π float64 = 3.14159

type SQUARE struct {
	length float64
}

type CIRCLE struct {
	radius float64
}

type SHAPE interface {
	area() float64
}

func main() {

	squ := SQUARE{
		20.00,
	}
	info(&squ)		//	Receiver is non-pointer, so value can be a pointer or a non-pointer.

	circle := CIRCLE{
		12.54,
	}

	info(&circle)	//	Receiver is non-pointer, so value must be a pointer.
					//	cannot use circle (variable of type CIRCLE) as SHAPE value in argument to info: CIRCLE does not implement SHAPE (method area has pointer receiver)

	fmt.Printf("%T\n%T\n", &circle, &squ)
}

func (cr *CIRCLE) area() float64 {
	return ((π * cr.radius) * (π * cr.radius))
}

func (sq SQUARE) area() float64 {
	return (sq.length * sq.length)
}

func info(sh SHAPE) {
	fmt.Println(sh.area())
	switch sh.(type) {
	case *CIRCLE:
		fmt.Println("Circle")
	case SQUARE:
		fmt.Println("Square")
	}
}
