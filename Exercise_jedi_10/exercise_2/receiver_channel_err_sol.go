package main

import (
	"fmt"
)

func main() {
  // below is initialization of receiver only channel
  // we cannot use it to sender channel
  // make cs as general/bidirectional channel

	//cr := make(<-chan int) //error point
  cr := make(chan int)     // solution

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
