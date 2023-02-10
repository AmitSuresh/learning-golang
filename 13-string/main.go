package main

import "fmt"

func main() {
	s:= `"Hello, 
	World."`

	fmt.Printf("%T, %v", s, s)
	bs := []byte(s)
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("\n%#U\t", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("\nindex\t%d, hex\t%#x\n", i, v)
	}
}