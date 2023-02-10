// in this package, we look into an efficient and inefficient way of doing the same function.
package efficiency

import "strings"

//InefficientWay takes in a slice of strings to be joined and returns the joined string in an inefficient way.
func InefficientWay(slices []string) string {
	str := slices[0]

	for _, word := range slices[1:] {
		str += " "
		str += word
	}
	return str
}

//EfficientWay takes in a slice of strings to be joined and returns the joined string in an Efficient way by utilizing a function in the standard package strings.Join([]string) string
func EfficientWay(slices []string) string {
	return strings.Join(slices, " ")
}
