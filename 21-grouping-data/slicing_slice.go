package main

import "fmt"

func main() {
	x := []int{25, 342, 52, 56, 7,34,64,43,2,76,76,6}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[1:3])
	fmt.Println(x[6:11])
	fmt.Println(x[1])
	fmt.Println(x[3])
	fmt.Println(x[6])
	fmt.Println(x[11])
	
	fmt.Println("\nbelow1")

	a := x[1:3]
	fmt.Println(a)
	b := x[6:11]
	fmt.Println(b)
	h := []int{}
	//h = append([]int{1}, b...)
	//h = append(a, b...)
	h = append(x[1:3], x[6:11]...)
	//h = append(x[1:3], b...)
	fmt.Printf("\nh = append(x[1:3], x[6:11]...)is: %v\n",h)

	fmt.Println("\nbelow2")

	y := append(x[1:3], x[6:11]...)
	fmt.Printf("\ny := append(x[1:3], x[6:11]...)is: %v\n",y)

	fmt.Println("\nbelow3")


	for i := 0; i <=3; i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println("\nbelow4")
	
	fmt.Println()
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	i := []int{}
	i = append(i, a...)
	fmt.Printf("\ni is: %v\n",i)
	i = append(i, b...)
	fmt.Printf("\ni is: %v\n",i)
}