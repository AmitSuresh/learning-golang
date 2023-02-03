package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //Cancel when we are finished.

	for n := range gen(ctx) {
		fmt.Println(n)
		time.Sleep(time.Millisecond * 10)
		if n == 5 {
			fmt.Println("breaking at 5")
			break
		}
		fmt.Println("after 5")
	}
	fmt.Println("outside 5")
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("After cancel, we get here at Done.")
				return
			case dst <- n:
				n++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	return dst
}
