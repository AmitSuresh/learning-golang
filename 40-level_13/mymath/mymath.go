// mymath package is used to facilitate CenteredAvg function.
// we are using this for level 13 exercises.
package mymath

import (
	"sort"
)

// CenteredAvg computes the average of a list of numbers
// after removing the largest and smallest values.
func CenteredAvg(xi []int) float64 {
	// Sorts the slice asc.
	sort.Ints(xi)

	// Iterates through second value to value before the last.
	xi = xi[1:(len(xi) - 1)]

	// Range through the slice, add up all the values and assign it to a variable.
	n := 0
	for _, v := range xi {
		n += v
	}

	// Divide the total value by the number of elements in slice.
	f := float64(n) / float64(len(xi))

	// Returns the average.
	return f
}
