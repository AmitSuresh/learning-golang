package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const num = 50

func main() {
	counter := 0
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			value := counter
			value++
			runtime.Gosched()
			counter = value
			fmt.Println("Counter", counter)
			fmt.Println("value", value)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
