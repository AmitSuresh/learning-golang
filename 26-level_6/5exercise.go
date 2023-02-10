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
	info(squ)

	circle := CIRCLE{
		12.54,
	}

	info(circle)
}

func (cr CIRCLE) area() float64 {
	return ((π * cr.radius) * (π * cr.radius))
}

func (sq SQUARE) area() float64 {
	return (sq.length * sq.length)
}

func info(sh SHAPE) {
	fmt.Println(sh.area())
	switch sh.(type) {
	case CIRCLE:
		fmt.Println("Circle")
	case SQUARE:
		fmt.Println("Square")
	}
}
