package main

import "fmt"

func main() {
	a := bar(foo, 2, 3)
	fmt.Println(a)

	em()
}

func foo(a, b int) int {
	return a + b
}

// here when we are using callback we have to create a funciton
// which takes functoin as it's one argument and other data arguments
// which are then used by the function
// like x and y are parameters for func bar() which uses function f
// to calculate sum of x and y
// so here we pass foo as argument to bar to calculate sum of x and y
func bar(f func(a, b int) int, x int, y int) int {
	return f(x, y)
}

// below we used and function with no parameters
func em() {
	fmt.Println("function em")
}

// in below function we passed a function as parameter with not parameters in it
// now we can pass em in me and can call em form me
func me(f func()) {
	f()
}
