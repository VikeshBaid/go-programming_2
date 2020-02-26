package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martini`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`kumar_gulshan`] = []string{`Bad Man`, `Villian`, `hindi movies`}

	for k, vs := range m {
		fmt.Printf("Key: %v\n", k)
		for i, v := range vs {
			fmt.Printf("\tIndex value: %v\t\t Value: %v\n", i, v)
		}
	}

  delete(m, `kumar_gulshan`)

  fmt.Println("\nAfter deletion of a key form map\n")

  for k, vs := range m {
		fmt.Printf("Key: %v\n", k)
		for i, v := range vs {
			fmt.Printf("\tIndex value: %v\t\t Value: %v\n", i, v)
		}
	}
}
