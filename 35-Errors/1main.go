package main

import "fmt"

func main() {
	f, err := fmt.Print("Glory")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("", f)
}
