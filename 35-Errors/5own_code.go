package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("names.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer copy_new_text(f)

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
func copy_new_text(filename *os.File) {
	new_string := strings.NewReader("\nAdding this to existing file.")
	_, err := new_string.WriteTo(filename)
	if err != nil {
		fmt.Println(err)
	}
	filename.Close()
}
