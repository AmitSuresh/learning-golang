package main

import "fmt"

func main() {
	fmt.Printf("Years alive:\n")
	x := 1997
	for {
		if x >= 2022 {
			break
		}
		fmt.Println(x)
		x++
	}
}