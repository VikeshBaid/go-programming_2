package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) speak() {
  // we have to use p.name & p.age instead of name and age
  // to let compiler know that speak wants name and age
  // value from person type
  // receiver are given to attach a method to a type so that
  // the method can use it's identifiers
  // it can be of pointer type and non pointer type
	fmt.Printf("My name is %s and I am %d years old\n", p.name, p.age)
}

type human interface {
  //
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "James Bond",
		age:  32,
	}
  // We have to pass the address of person type variable because
  // the speak method which is attached to a pointer type person
  // expects the address of person type identifier
	saySomething(&p1) // this will work
  //saySomething(p1) // this won't work
}
