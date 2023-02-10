package main

import "fmt"

func main() {
	bond_james := []string{`Shaken, not stirred`, `Martinis`, `Women`}
	moneypenny_miss := []string{`James Bond`, `Literature`, `Computer Science`}
	no_dr := []string{`Being evil`, `Ice cream`, `Sunsets`}

	y := map[string][]string{
		`bond_james`:      bond_james,
		`moneypenny_miss`: moneypenny_miss,
		`no_dr`:           no_dr,
	}

	fmt.Println("length of y:", len(y))
	y["mr_killy"] = []string{"Killing", "Hiking", "Shooting"}

	fmt.Println("length of y:", len(y))
	if val, okk := y[`no_dr`]; okk {
		fmt.Println("value: ", val)
		fmt.Println("value in map: ", y["no_dr"])
		delete(y, `no_dr`)
		fmt.Println("value: ", val)
		fmt.Println("value in map: ", y["no_dr"])
	}
	fmt.Println("length of y:", len(y))

	for key_string_of_map, value_slice_of_map := range y {
		fmt.Printf("\nkey_string_of_map: %v\nvalue_slice_of_map:%v\n\n", key_string_of_map, value_slice_of_map)
		for string_number, string_value := range value_slice_of_map {
			fmt.Printf("string_number: %v, string_value: %v\n\n", string_number, string_value)
		}
	}

}
