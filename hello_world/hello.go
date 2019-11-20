package main

import (
	"fmt"
)

func main() {
  // description of n and err at https://godoc.org/fmt
	n, err := fmt.Println("Hello Go Programming language", 42, true)
  fmt.Println(n)
  fmt.Println(err)
}
