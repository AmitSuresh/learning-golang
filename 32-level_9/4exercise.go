package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mx sync.Mutex

const num = 50

func main() {
	counter := 0
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			mx.Lock()
			value := counter
			value++
			time.Sleep(time.Millisecond*num)
			counter = value
			fmt.Println("Counter", counter)
			fmt.Println("value", value)
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
