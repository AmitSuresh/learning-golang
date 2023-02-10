package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

const num = 50

func main() {
	var counter int64
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter: ", atomic.LoadInt64(&counter))
			time.Sleep(time.Millisecond * num)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
