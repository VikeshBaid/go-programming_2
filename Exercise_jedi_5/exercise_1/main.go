package main

import (
	"fmt"
)

type person struct {
	first        string
	last         string
	fav_icecream []string
}

// func (p person) String() string {
// 	return fmt.Sprintf("%s %s loves %v", p.first, p.last, p.fav_icecream)
// }

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

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.fav_icecream {
		fmt.Printf("\t%v\t%v\n", i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.fav_icecream {
		fmt.Printf("\t%v\t%v\n", i, v)
	}
}
