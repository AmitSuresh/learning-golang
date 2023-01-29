package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	add_value := make(chan int)
	sub_value := make(chan int)

	go func(a int, b int) {
		time.Sleep(time.Second * 11)
		fmt.Println("Inside 1")
		add_value <- a + b
	}(23, 43)

	go func(a int, b int) {
		time.Sleep(time.Microsecond)
		fmt.Println("Inside 2")
		sub_value <- a - b
	}(43, 12)

	for i := 0; i < 2; i++ {
		// Await both of these values
		// simultaneously, printing each one as it arrives.
		select {
		case msg1 := <-add_value:
			fmt.Println("received1", msg1)
		case msg2 := <-sub_value:
			fmt.Println("received2", msg2)
		}
	}

	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
}
