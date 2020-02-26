package main

//race condition resolve using sync/atomic
import (
	"fmt"
	// "runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var counter int64
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter :", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		// fmt.Println("GOroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	// fmt.Println("GOroutines:", runtime.NumGoroutine())
	fmt.Println("count: ", counter)
}
