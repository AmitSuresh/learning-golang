// I used speak package is used to learn benchmarking.
package speak

import "fmt"

//Hello function returns a string.
//It takes in a string as a argument and returns a string greeting the string parameter.
func Hello(str string) string {
	return fmt.Sprint("Greetings..", str)
}
