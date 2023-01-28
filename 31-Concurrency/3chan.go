package main

import "fmt"

func dosome(x int) int {
	return x * 5
}
func main() {
	ch := make(chan int)
	go func() {
		ch <- dosome(5)
	}()
	fmt.Println(<-ch)
}
