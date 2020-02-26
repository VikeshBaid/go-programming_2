package main

import (
	"fmt"
  "github.com/go-programming_2/Exercise_jedi_13/exercise_1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(2),
	}

	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
