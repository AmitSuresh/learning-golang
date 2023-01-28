package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	fmt.Println("runtime CPUs: ", runtime.NumCPU())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	//fmt.Println("time: ", time.Duration(gs)*(10*time.Microsecond))
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Millisecond * (time.Duration(v)))
			fmt.Println("time: ", time.Millisecond*(time.Duration(v)))
			runtime.Gosched()
			v++
			fmt.Println("v: ", v)
			counter = v
			fmt.Println("counter: ", counter)
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("counter: %v", counter)

}
