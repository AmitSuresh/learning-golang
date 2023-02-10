package main

import "fmt"

func main() {
	js := []string{"Johnny", "Silverhand", "Chocolate", "Sprinkled Water"}
	fmt.Println(js)
	v := []string{"-", "V", "Strawberry", "Water"}
	fmt.Println(v)

	mp := [][]string{js, v}
	fmt.Println(mp)
}
