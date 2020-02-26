package main

import (
	"fmt"
)

func main() {
  // below is initialization of sender only channel
  // we cannot use it to receiver channel
  // make cs as general/bidirectional channel

  //cs := make(chan<- int)  //error point
  cs := make(chan int)      // solution

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
