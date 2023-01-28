package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	fmt.Println("runtime CPUs: ", runtime.NumCPU())

	var counter int64 = 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	//fmt.Println("time: ", time.Duration(gs)*(10*time.Microsecond))
	for i := 0; i < gs; i++ {
		go func() {
			time.Sleep(time.Millisecond * (time.Duration(atomic.LoadInt64(&counter))))
			fmt.Println("time: ", time.Millisecond*(time.Duration(atomic.LoadInt64(&counter))), "counter: ", atomic.LoadInt64(&counter))
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter from atomic: ", atomic.LoadInt64(&counter))
			runtime.Gosched()
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("counter: %v", counter)

}
