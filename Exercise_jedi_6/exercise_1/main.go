package main

import "fmt"

func main() {
	a := foo()
	x, y := boo()

	fmt.Println(a, x, y)
}

func foo() int {
	return 42
}

func boo() (int, string) {
	return 24, "hello"
}
