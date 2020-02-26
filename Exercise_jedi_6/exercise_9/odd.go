package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	odds := odd(sum, ii...)
	fmt.Println("Odd number sum b/w 0 to 9: ", odds)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func odd(f func(xi ...int) int, yi ...int) int {
	var vi []int
	for _, v := range yi {
		if v%2 != 0 {
			vi = append(vi, v)
		}
	}
	return f(vi...)
}
