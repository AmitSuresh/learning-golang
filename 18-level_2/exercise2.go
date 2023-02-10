package main

import (
	"fmt"
)

func main() {
	g := 10 == 20
	h := 10 <= 20
	i := 10 >= 20
	j := 10 != 20
	k := 10 < 20
	l := 10 > 20
	
	fmt.Println(g,h,i,j,k,l)
}