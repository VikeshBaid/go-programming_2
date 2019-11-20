package main

import (
	"fmt"
)

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE STRING
var z = "var can be used to declare a variable outside the func main() but short hand operator can't do that."

func main() {
	x := 42 // short hand operator, declaing and assigning

	// DECLARE the variable "y"
	// ASSIGN the value 43
	// declare & assign = initilization
	var y = 42 // var keyword, using to declare and assign a variable
	fmt.Println("using short hand printing x", x)
	fmt.Println("using var keyword printing y inside the func main scope", y)
	fmt.Println("using var keyword printing z which is declared ouside func main.", z)
}
