package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go foo(c, 10)

	bar(c)

}

func foo(c chan<- int, limit int) {
	for i:=0;i<limit;i++{
		c <- 7
	}
	close(c)
}

func bar(c <-chan int) {
	// Await both of these values
	// simultaneously, printing each one as it arrives.
	/* select {
	case s := <-c:
		fmt.Println("c is: ", s)
	} */
	for s := range c{
		fmt.Println(s)
	}
	/* fmt.Println("c is: ", <-c) */
}
