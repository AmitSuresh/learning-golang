package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[j].Age < a[i].Age }

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name > a[j].Name }

func main() {
	num := []int{2, 5, 3, 7, 1, 7, 9, 4, 1}
	sort.Ints(num)
	fmt.Println(num)
	strs := []string{"monkey", "lion", "tiger", "gorilla", "zebra", "elephant"}
	sort.Strings(strs)
	fmt.Println(strs)

	npc1 := []person{
		{"Jackie",
			55},
		{"V",
			34},
		{"Johnny",
			51},
		{"Misty",
			36},
	}

	fmt.Printf("\nThe slice: %v\n", npc1)
	sort.Sort(ByAge(npc1))
	fmt.Printf("\nThe slice after sorting by age: %v\n", npc1)

	sort.Sort(ByName(npc1))
	fmt.Printf("\nThe slice after sorting by name: %v\n", npc1)

}
