package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	wg.Add(2)
	go func() {
		fmt.Println("First go")
		fmt.Println("GoRoutine1: ", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {
		fmt.Println("Second go")
		fmt.Println("GoRoutine2: ", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
}
