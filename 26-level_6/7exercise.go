package main

import (
	"fmt"
)

func main() {
	func_to_variable := func() {
		fmt.Println("Assigned func to variable and called it.")
	}
	func_to_variable()
}
