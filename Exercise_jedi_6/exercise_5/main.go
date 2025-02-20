package main

import (
  "fmt"
  "math"
)

type square struct {
   length float64
}

type circle struct {
  radius float64
}

func (c circle) area() float64  {
  return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
  return s.length * s.length
}

type shape interface {
  area() float64
}

func info(s shape)  {
  fmt.Println(s.area())
}

func main()  {
  cir := circle {
    radius: 4,
  }

  squ := square {
    length: 32,
  }

  info(cir)
  info(squ)
}
