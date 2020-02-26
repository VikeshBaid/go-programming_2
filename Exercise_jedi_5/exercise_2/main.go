package main

import (
	"fmt"
)

type person struct {
	first        string
	last         string
	fav_icecream []string
}

func main() {
	p1 := person{
		first:        "James",
		last:         "Bond",
		fav_icecream: []string{"Martini", "Belgium CHocolate"},
	}

	p2 := person{
		first:        "Miss",
		last:         "Moneypenny",
		fav_icecream: []string{"Bubblegum", "Chocochips"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.fav_icecream {
			fmt.Println(i, val)
		}
		fmt.Println("----------")
	}
}
