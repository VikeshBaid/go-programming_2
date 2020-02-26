package main

import (
  "fmt"
)

type Vector int
var x Vector

func main() {
  x = 43
  fmt.Println(x)
  fmt.Printf("%T\n", x)
  x = 42
  fmt.Println(x)
}
