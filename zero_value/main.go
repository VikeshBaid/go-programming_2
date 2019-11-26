package main

import (
	"fmt"
)
//use var for
// zero value and
// package scope

var x string
var y int
var z float32

func main() {
	// Here we are goin to see ZERO VALUE for some types
	// more detail on https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.vgdf3swpywdi
	// above we have declared some variable with there TYPES

	fmt.Println("zero value of a string type is ", x, "here.")
	fmt.Printf("%T\n", x)

	fmt.Println("zero value of a int type is ", y, "here.")
	fmt.Printf("%T\n", y)

	fmt.Println("zero value of a float type is ", z, "here.")
	fmt.Printf("%T\n", z)
}
