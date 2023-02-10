package main

import (
	"fmt"
	"runtime"
)

func main() {
	outerchannel := make(chan int)
	fmt.Println("numgo:", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func(channel_1 chan int) {
			for j := 0; j < 10; j++ {
				channel_1 <- j
				fmt.Println("inside index j", j)
			}
			fmt.Println("numgo:", runtime.NumGoroutine())
		}(outerchannel)
	}
	for i := 0; i < 100; i++ {
		fmt.Println("asdf", <-outerchannel)
	}
}

