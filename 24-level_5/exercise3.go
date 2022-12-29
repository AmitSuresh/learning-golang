package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	my_truck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "White",
		},
		fourWheel: true,
	}

	my_sedan := sedan{
		vehicle{
			4, "Black",
		},
		true,
	}

	fmt.Printf("Car Information:\n")
	fmt.Printf("\tMy Truck:\n\t\tDoors:%v\n\t\tColor:%v\n\t\tFour Wheel Drive:%v\n", my_truck.doors, my_truck.color, my_truck.fourWheel)
	fmt.Printf("\tMy Sedan:\n\t\tDoors:%v\n\t\tColor:%v\n\t\tLuxury:%v\n", my_sedan.doors, my_sedan.color, my_sedan.luxury)
}
