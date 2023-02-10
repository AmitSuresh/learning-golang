package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("func exp.")
	}
	f()

	add_2 := func(a, b int) {
		fmt.Println("Sum:", a+b)
	}
	add_2(5, 6)
	add_slice := func(a ...int) {
		sum := 0
		for i := range a {
			sum += a[i]
		}
		fmt.Println("1add_slice:", sum)
		fmt.Printf("2add_slice:%v\n\n", a)
	}
	a := []int{2, 3, 4}
	add_slice(a...)

	add_slice([]int{3, 4, 5}...)
	add_slice()
	add_slice(2,3,4)
}
