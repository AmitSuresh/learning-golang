package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (bA ByAge) Len() int           { return len(bA) }
func (bA ByAge) Less(i, j int) bool { return bA[i].Age < bA[j].Age }
func (bA ByAge) Swap(i, j int)      { bA[i], bA[j] = bA[j], bA[i] }

type ByLast []user

func (bL ByLast) Len() int           { return len(bL) }
func (bL ByLast) Less(i, j int) bool { return bL[i].Last < bL[j].Last }
func (bL ByLast) Swap(i, j int)      { bL[i], bL[j] = bL[j], bL[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("\nRaw below:")

	for _, usr := range users {
		fmt.Println("User:\t", usr.First, usr.Last, usr.Age)
		sort.Strings(usr.Sayings)
		for _, saying := range usr.Sayings {
			fmt.Println("Sayings: ", saying)
		}
	}

	sort.Stable(ByAge(users))

	fmt.Println("\nSort by Age below:")

	for _, usr := range users {
		fmt.Println("User:\t", usr.First, usr.Last, usr.Age)
		for _, saying := range usr.Sayings {
			fmt.Println("Sayings: ", saying)
		}
	}

	sort.Stable(ByLast(users))

	fmt.Println("\nSort by Last below:")

	for _, usr := range users {
		fmt.Println("User:\t", usr.First, usr.Last, usr.Age)
		for _, saying := range usr.Sayings {
			fmt.Println("Sayings: ", saying)
		}
	}

}
