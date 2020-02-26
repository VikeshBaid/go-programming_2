package main

import (
	"fmt"
)

const a int = 32
const b = 32.34

func main() {
	fmt.Println(a)
	fmt.Println(b)
  fmt.Printf("%T\n", a)
  fmt.Printf("%T\n", b)
}
