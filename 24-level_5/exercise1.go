package main

import "fmt"

type person struct {
	first_name, last_name string
	favorite_flavors      []string
}

func main() {

	me := person{
		"Amit",
		"Suresh",
		[]string{
			"Chocolate",
			"Vanilla",
		},
	}

	var my_fav_flavors []string
	my_fav_flavors = append(my_fav_flavors, me.favorite_flavors...)

	friend_fav_flavors := []string{"Butterscotch", "other"}

	friend := person{
		first_name:       "Jackie",
		last_name:        "Welles",
		favorite_flavors: friend_fav_flavors,
	}

	fmt.Printf("I am %v %v.\n", me.first_name, me.last_name)
	fmt.Printf("I like:\n")
	for i := range me.favorite_flavors {
		fmt.Printf("\t%v\n", me.favorite_flavors[i])
	}

	fmt.Printf("I am %v %v.\n", friend.first_name, friend.last_name)
	fmt.Printf("I like:\n")
	for i := range friend.favorite_flavors {
		fmt.Printf("\t%v\n", friend.favorite_flavors[i])
	}

	fmt.Printf("\n\n%v\n\n", my_fav_flavors)
	fmt.Printf("\n\n%v", friend_fav_flavors)
}
