package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)

	go func() {
		c <- 42
  }()

  v, ok := <- c
	fmt.Println("Before close",v, ok)

	close(c)

	v, ok = <- c
	fmt.Println("After close",v, ok)

}
