package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	fmt.Println()
	p1 := person{
		"Johnny",
		"Silverhand",
		52,
	}
	p2 := person{
		"sdf", "asdf", 42,
	}
	pess := []person{p1, p2}
	fmt.Println(pess)
	jsn, err := json.Marshal(pess)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsn))
	adsf := make([]person, 0, 0)
	err1 := json.Unmarshal(jsn, &adsf)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(adsf)
}
