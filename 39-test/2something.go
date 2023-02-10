package main

import (
	"fmt"

	"github.com/AmitSuresh/learning-golang/39-test/something"
)

func main() {
	slices := []int{10, 10}
	sum_of_num := something.Sum(slices...)
	fmt.Println(sum_of_num)
}
