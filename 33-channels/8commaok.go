// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
)

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func main() {
	//var seconds1 int
	//var ok1 bool
	//seconds1, ok1 = timeZone["ST"]
	//fmt.Println(seconds1)
	//fmt.Println(ok1)

	some := offset("MST")
	fmt.Println(some)

}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}
