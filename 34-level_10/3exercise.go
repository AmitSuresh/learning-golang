package main

import (
	"fmt"
)

func main() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
			close(c)
		}()
		return c
	}()
	func(c <-chan int) {
		for n := range c {
			fmt.Println(n)
		}
	}(c)

	fmt.Println("about to exit")
}
