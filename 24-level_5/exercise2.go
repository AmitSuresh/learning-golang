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

	map_people := map[string]person{
		me.last_name:     me,
		friend.last_name: friend,
	}

	fmt.Printf("Map of people: %v\n", map_people["Suresh"])
	fmt.Printf("Map of people: %v", map_people["Welles"])

	for k, v := range map_people {
		fmt.Printf("\n\n\nName:%v\n", k)
		for key, value := range v.favorite_flavors {
			fmt.Printf("\t%v. %v\n", key, value)
		}

	}

	for i, v := range map_people[me.last_name].favorite_flavors {
		fmt.Printf("\n%v. value: %v", i, v)
	}

	for i, v := range map_people[friend.last_name].favorite_flavors {
		fmt.Printf("\n%v. value: %v", i, v)
	}
}
