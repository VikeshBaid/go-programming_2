package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
	circumference() float64
}
// in function definition the receiver is of pointer type
// the argument pass to it should be using "&" i.e address of the value
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
	fmt.Println("circumference", s.circumference())
}

func main() {
	c1 := circle{5}
	info(&c1)
}
