package main

import (
  "fmt"
  "github.com/go-programming_2/Exercise_jedi_12/exercise_1/dog"
)

type canine struct {
  name string
  age int
}
func main() {
  fido := canine{
    name : "Fido",
    age : dog.Years(2),
  }

  fmt.Println(fido)
}
