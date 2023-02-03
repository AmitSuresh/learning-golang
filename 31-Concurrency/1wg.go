package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg aitGroup
var wg1 sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	wg.Add(3)
	wg1.Add(2)
	go zar()
	go foo()
	bar()
	jar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	wg.Wait()

	go far()
	go too()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	wg1.Wait()
}

func foo() {
	for i := 0; i < 11; i++ {
		fmt.Printf("foo:%v\t", i)
	}
	fmt.Println()
	wg.Done()
	go doo()
}

func bar() {
	for i := 0; i < 11; i++ {
		fmt.Printf("bar:%v\t", i)
	}
	fmt.Println()
}

func doo() {
	for i := 0; i < 11; i++ {
		fmt.Printf("doo:%v\t", i)
	}
	fmt.Println()
	wg.Done()
}

func jar() {
	for i := 0; i < 11; i++ {
		fmt.Printf("jar:%v\t", i)
	}
	fmt.Println()
}

func zar() {
	for i := 0; i < 1234; i++ {
		fmt.Printf("zar:%v\t", i)
	}
	fmt.Println()
	wg.Done()
}

func too() {
	for i := 0; i < 11; i++ {
		fmt.Printf("too:%v\t", i)
	}
	fmt.Println()
	wg1.Done()
}

func far() {
	for i := 0; i < 1234; i++ {
		fmt.Printf("far:%v\t", i)
	}
	fmt.Println()
	//wg1.Done()
	go ear()
}

func ear() {
	for i := 0; i < 11; i++ {
		fmt.Printf("ear:%v\t", i)
	}
	fmt.Println()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	wg1.Done()
}
