package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
	}{
		first: "Todd",
		last:  "Mcleod",
	}

	fmt.Println(p1.first, p1.last)
}
