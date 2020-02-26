package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		// sending data to channel( channel type identifier)
		c <- 42
	}()

	// receiving data from channel( channel type identifer)
	fmt.Println(<-c)

	// to use channel in Golang the sender and receiver should be execute
	// simultaneously

}
