package main

import "fmt"

func main() {
	s := "911"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n%b\n%x\n%d\n%#X", n, n, n, n, n)
}