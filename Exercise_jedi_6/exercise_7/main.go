package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("assigned function")
	}

	a()

  g := a
	fmt.Printf("%T\n", g)
}
