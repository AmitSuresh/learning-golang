package main

import (
	"fmt"
)

func main() {
	x, y := 10, 10
	readvalue(getvalues(x, y))
}

func readvalue(channel_1 chan int) {
	/* for {
		select {
		case i := <-channel_1:
			fmt.Println("value of i is: ", i)
		default:
			fmt.Println()
			return
		}
	} */
	
	for i := 0; i < 100; i++ {
		fmt.Println("asdf", <-channel_1)
	}
}

func getvalues(x, y int) chan int {
	channel_1 := make(chan int)
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