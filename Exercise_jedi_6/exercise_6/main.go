package main

import "fmt"

func main() {
	func() {
		a, b := 3, 6
		fmt.Println(a + b)
	}()
}
