package main

import "fmt"

func main() {
	bond_james := []string{`Shaken, not stirred`, `Martinis`, `Women`}
	moneypenny_miss := []string{`James Bond`, `Literature`, `Computer Science`}
	no_dr := []string{`Being evil`, `Ice cream`, `Sunsets`}

	y := map[string][]string{
		`bond_james`: bond_james,
		`moneypenny_miss`: moneypenny_miss,
		`no_dr`: no_dr,
	}
	
	y["mr_killy"] = []string{"Killing","Hiking", "Shooting"}

	for key_string_of_map, value_slice_of_map := range y {
		fmt.Printf("\nkey_string_of_map: %v\nvalue_slice_of_map:%v\n\n", key_string_of_map, value_slice_of_map)
		for string_number, string_value := range value_slice_of_map {
			fmt.Printf("string_number: %v, string_value: %v\n\n", string_number, string_value)
		}
	}
	z := map[string]map[string][]string{
		`map1`: y,
	}
	fmt.Println(z)

	for key_string_of_map, value_slice_of_map := range z {
		fmt.Printf("\nkey_string_of_map: %v\nvalue_slice_of_map:%v\n\n", key_string_of_map, value_slice_of_map)
		for string_number, string_value := range value_slice_of_map {
			fmt.Printf("string_number: %v, string_value: %v\n\n", string_number, string_value)
		}
	}
	a := map[string]map[string]map[string][]string{
		`outermap1`: z,
	}
	fmt.Println(a)

	for key_string_of_map, value_slice_of_map := range a {
		fmt.Printf("\nkey_string_of_map: %v\nvalue_slice_of_map:%v\n\n", key_string_of_map, value_slice_of_map)
		for string_number, string_value := range value_slice_of_map {
			fmt.Printf("string_number: %v, string_value: %v\n\n", string_number, string_value)
			for string_number, string_value := range string_value {
				fmt.Printf("string_number: %v, string_value: %v\n\n", string_number, string_value)
			}
		}
	}

}