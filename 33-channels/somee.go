package main

import (
	"fmt"
)

func main() {
	x, y := 10, 10
	channel_1, q_1 := getvalues(x, y)
	readvalue(channel_1, q_1)
}

func readvalue(channel_1, q_1 chan int) {
	for {
		select {
		case i := <-channel_1:
			fmt.Println("value of i is: ", i)
		}
	}
}

func getvalues(x, y int) <-chan int, <-chan int{
	channel_1 := make(chan int)
	q_1 := make(chan int)
	for i := 0; i < x; i++ {
		go func(channel_1 chan int) {
			for i := 0; i < y; i++ {
				channel_1 <- i
				fmt.Println("inside index ", i)
			}
		}(channel_1)
	}
	return channel_1
}