package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello world")

}

func foo() {
  defer func() {
    fmt.Println("foo Defer Ran")
  }()
	fmt.Println("foo Ran")
}
