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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
	fmt.Println("circumference", s.circumference())
}

func main() {
	c := circle{5}
	info(c)
}
