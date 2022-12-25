package main

import "fmt"

func main() {
	favSport := "Muay Thai"
	switch favSport {
	case "running":
		fmt.Printf("running")
	case "MMA":
		fmt.Printf("MMA")
	case "Muay Thai", "Kickboxing":
		fmt.Printf("Muay Thai, Kickboxing")
	case "Soccer":
		fmt.Printf("Soccer")
	default:
		fmt.Printf("Muay Thai by default.")
	}
}
