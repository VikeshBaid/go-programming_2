package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-ch)
	}

	fmt.Println("about to exit")
}
