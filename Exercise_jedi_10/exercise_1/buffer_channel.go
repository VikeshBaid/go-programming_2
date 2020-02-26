package main

import (
	"fmt"
)

func main() {
	// below "1" is the buffer size
	// buffer stores the data which is send to a channel
	// then send it back when data is received by channel
	// we have o give a particular size to buffer
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
