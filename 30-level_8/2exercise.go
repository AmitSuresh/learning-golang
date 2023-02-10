package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
	Sayings     []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println("\n", s)
	fmt.Println()

	unjson1 := make([]person, 0, 0)

	err := json.Unmarshal([]byte(s), &unjson1)
	if err != nil {
		fmt.Println("error is: ", err)
		return
	}
	fmt.Println("\n", unjson1)
	fmt.Println()

	for _, person := range unjson1 {
		fmt.Println("Person:\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\tSaying:\t", saying)
		}
	}

}
