package main

import (
	"fmt"
	"runtime"
)

var x rune
var y float64

func main() {
	x = 127
	y = 23.999
	fmt.Printf("value and type:\n\t%T\t%v\n\t%T\t%v\n",x,x,y,y)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}