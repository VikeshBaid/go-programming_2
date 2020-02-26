package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("main goroutine")

	fmt.Println("Num of CPUs", runtime.NumCPU())
	fmt.Println("Num Goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("First goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second goroutine")
		wg.Done()
	}()
	fmt.Println("Num of CPUs mid", runtime.NumCPU())
	fmt.Println("Num Goroutine mid", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Num of CPUs end", runtime.NumCPU())
	fmt.Println("Num Goroutine end", runtime.NumGoroutine())
	fmt.Println("main goroutine ends")

}
