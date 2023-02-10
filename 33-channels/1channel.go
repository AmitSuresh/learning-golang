package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go foo(c)

	bar(c)

}

func foo(c chan<- int) {
	c <- 7
}

func bar(c <-chan int) {
	// Await both of these values
	// simultaneously, printing each one as it arrives.
	select {
	case s := <-c:
		fmt.Println("c is: ", s)
	}
	/* fmt.Println("c is: ", <-c) */
}
