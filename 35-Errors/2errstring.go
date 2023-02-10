package main

import (
	"fmt"
)

func main() {

	err := new("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println()
		fmt.Print(err)
		fmt.Println()
		fmt.Printf("\n\n%T\n\n", err)
		fmt.Println(&err)
		fmt.Println()
	}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func new(text string) *errorString {

	ff := &errorString{text}
	fmt.Printf("ff is %T, %v\n", ff, ff)
	fmt.Println()
	fg := errorString{text}
	fmt.Printf("fg is %T, %v\n", fg, fg)
	fmt.Println(fg)
	return ff
}
