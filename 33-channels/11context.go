package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check 1:", ctx.Err())
	fmt.Println("num goroutine 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("jumping here and returning.")
				return
			default:
				n++
				time.Sleep(time.Millisecond*100)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Error check 2:", ctx.Err())
	fmt.Println("num goroutine 2:", runtime.NumGoroutine())

	fmt.Println("About to cancel context.")
	cancel()
	fmt.Println("Canceled")

	time.Sleep(time.Second)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutine 3:", runtime.NumGoroutine())
}
