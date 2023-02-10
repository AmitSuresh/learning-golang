// something package is used to add a slice of ints.
package something

// Sum from the package something returns sum of slice of ints.
func Sum(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}
