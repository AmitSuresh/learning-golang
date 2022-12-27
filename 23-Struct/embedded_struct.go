package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

type character struct {
	person
	game string
	age  int
}

func main() {

	cyb := character{
		person: person{
			first_name: "",
			last_name:  "Welles",
			age:        35,
		},
		game: "cyberpunk2077",
		age:  3,
	}

	person1 := person{
		first_name: "Johnny",
		last_name:  "Silverhand",
		age:        45,
	}

	person2 := person{
		first_name: "",
		last_name:  "V",
	}

	person3 := person{
		first_name: "Geralt",
		last_name:  "Of Rivia",
	}

	cyb1 := character{
		person: person1,
		game:   "cyberpunk2077",
	}
	cyb2 := character{
		person: person2,
		game:   "cyberpunk2077",
	}
	wit := character{
		person: person3,
		game:   "witcher",
	}

	fmt.Printf("Person1 is: %v\nPerson2 is: %v\n", person1, person2)
	fmt.Printf("\nMr %v is %v-years old.", person1.last_name, person1.age)
	fmt.Printf("\nMr %v is %v-years old.", person2.last_name, person2.age)
	fmt.Printf("\nMr %v is %v-years old. The game %v is %v-years old.", cyb.last_name, cyb.person.age, cyb.game, cyb.age)	// when the inner type and outer type share the same key, outer one takes priority.
	fmt.Printf("\nMr %v of %v is %v-years old.", cyb1.first_name, cyb1.game, cyb1.age)
	fmt.Printf("\nMr %v of %v is %v-years old.", cyb2.first_name, cyb2.game, cyb2.age)
	fmt.Printf("\nMr %v of %v is %v-years old.", wit.first_name, wit.game, wit.age)
}
