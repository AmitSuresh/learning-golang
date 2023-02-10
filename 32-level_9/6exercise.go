package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOOS", runtime.GOOS)
	fmt.Println("GOROOT", runtime.GOROOT())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	fmt.Println("GOARCH", runtime.GOARCH)
	fmt.Println("GOROOT", runtime.GOROOT())
}
