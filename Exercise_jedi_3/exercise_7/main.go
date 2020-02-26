package main

import (
	"fmt"
)

func main() {
	i := 1
	if i == 3 {
		fmt.Println("if in action")
	} else if i == 2 {
		fmt.Println("i is 2")
	} else {
		fmt.Println("i is not 2 or 3")
	}
}
