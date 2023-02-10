package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

func main() {
	person1 := person{
		first_name: "Johnny",
		last_name:  "Silverhand",
		age:        45,
	}

	person2 := person{
		first_name: "",
		last_name:  "V",
	}

	fmt.Printf("Person1 is: %v\nPerson2 is: %v\n", person1, person2)
	fmt.Printf("\nMr %v is %v-years old.", person1.last_name, person1.age)
	fmt.Printf("\nMr %v is %v-years old.", person2.last_name, person2.age)
}
