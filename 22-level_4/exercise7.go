package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hellooo, James."}
	z := [][]string{x, y}
	fmt.Printf("Slice of a Slice: %v\n\n", z)

	for index_of_slice, slice := range z {
		fmt.Printf("\nindex_of_slice: %v, slice: %v\n\n", index_of_slice, slice)
		for index_of_string, string_in_slice := range slice {
			fmt.Printf("inside: index_of_string_in_slice: %v, string_in_slice: %v\n", index_of_string, string_in_slice)
		}
	}
}
