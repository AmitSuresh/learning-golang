package main

import (
	"fmt"

	"github.com/AmitSuresh/learning-golang/40-level_13/mymath"
)

func main() {
	slice := []int{31, 123, 43, 634, 12, 9, 6, 5, 5, 3, 12, 43}
	fmt.Println(mymath.CenteredAvg(slice))
	slice = []int{70, 70, 50, 20, 10, 10, 60, 70, 20, 40, 50, 80, 10, 50, 90, 100}
	fmt.Println(mymath.CenteredAvg(slice))
	slices := gen()
	for _, v := range slices {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
