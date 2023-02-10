package main

import (
	"fmt"
)

func main() {

	chann1, q := func() (<-chan int, <-chan int) {
		ch := make(chan int)
		q := make(chan int)
		go func() {
			for i := 0; i < 100; i++ {
				ch <- i
			}
			close(ch)
			close(q)
		}()
		return ch, q
	}()

	func(ch, q <-chan int) {
		for {
			select {
			case v := <-ch:
				fmt.Println("reading from switch statement: ", v)
			case <-q:
				fmt.Println("exiting from the quit channel",)
				return
			}
		}
	}(chann1, q)

	fmt.Println("about to exit")
}
