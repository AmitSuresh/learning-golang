package main

import "fmt"

func main() {
	x := struct {
		who        string
		angry      bool
		reason     []string
		skillLevel map[string]int
	}{
		"Amit",
		true,
		[]string{
			"Taking too long",
			"Talking too slow",
		},
		map[string]int{
			"Python": 9,
			"Go":     8,
		},
	}

	fmt.Printf("Who is angry?\n\t%v\nIs he angry now?\n\t%v\nWhy is he angry?\n", x.who, x.angry)
	for k, v := range x.reason {
		fmt.Printf("\t%v. %v\n", k, v)
	}
	fmt.Println("What are his skills?")
	for k, v := range x.skillLevel {
		fmt.Printf("\t%v. %v\n", k, v)
	}
}
