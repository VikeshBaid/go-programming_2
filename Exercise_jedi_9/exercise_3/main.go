package main

//creating race condition
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v1 := counter
			fmt.Println("v1 counter:", v1)
			v1++
			counter = v1
			fmt.Println("v1 counter after inc: ", v1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count: ", counter)
}
