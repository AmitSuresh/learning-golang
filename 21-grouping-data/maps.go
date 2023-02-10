package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	mp := map[string]int{
		"James Bond":      54,
		"Miss Moneypenny": 23,
	}
	fmt.Println(mp)

	fmt.Println(mp["James Bond"])

	fmt.Println(mp["Out Of Bounds"])

	mp["Johnny"] = 46	// this adds new record to the map.

	for k, v := range mp {
		fmt.Println(k, v)
	}

	v1, okk1 := mp["James Bond"]
	fmt.Println(v1)
	fmt.Println(okk1)

	v, ok := mp["Barry Allen"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := mp["Barry Allen"]; ok {
		fmt.Printf("Value 2is: %v", v)
		fmt.Println()
		fmt.Printf("Value 1is: %v", ok)
	}
	fmt.Println()
	if v, ok := mp["James Bond"]; ok {
		fmt.Printf("Value 2is: %v", v)
		fmt.Println()
		fmt.Printf("Value 2is: %v", ok)
	}

	fmt.Printf("\n\n\nAfter line\n>>>>>>>>>>>>>>>>>>>>>>>45<<<<<<<<<<<<<<<<<\n\n\n")

	del_mp := map[string]int{
		"James Bond":      54,
		"Miss Moneypenny": 23,
	}

	del_mp["Johnny"] = 46
	del_mp["Tim"] = 12
	del_mp["Sorted?"] = 112
	del_mp["ByKey!"] = 1192
	fmt.Println(del_mp)
	delete(del_mp, "Johnny")
	fmt.Println(del_mp)
	delete(del_mp, "empty")
	fmt.Println(del_mp)

	if val, okk := del_mp["James Bond"]; okk {
		fmt.Println("value: ", val)
		delete(del_mp, "James Bond")
	}
	if val1, okk1 := del_mp["ahdffsgh"]; okk1 {
		fmt.Println("value: ", val1)
		delete(del_mp, "ahdffsgh")
	}
	fmt.Println(del_mp)
}
