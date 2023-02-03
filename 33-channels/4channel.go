package main

import (
	"fmt"
)

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	//send
	go send(odd, even, quit)

	//receive
	receive(odd, even, quit)
}

func receive(o, e, q <-chan int) {
	for {
		select {
		case v := <-o:
			fmt.Println("from the odd channel: ", v)
		case v := <-e:
			fmt.Println("from the even channel: ", v)
		case <-q:
			fmt.Println("exiting from the quit channel: ")
			return
		}
	}
}

func send(o, e, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
