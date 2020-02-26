package main

import (
	"fmt"
)

func main() {
	a := []string{"James", "Bond", "Shaken not stirred"}
	b := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	xxs := [][]string{a, b}

	for i, xv := range xxs {
		fmt.Printf("index %v\n", i)
		for j, v := range xv {
			fmt.Printf("\tIndex Position: %v\t\t Value: %v\n", j, v)
		}
	}
}
