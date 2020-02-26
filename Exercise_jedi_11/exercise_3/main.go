package main

import (
	"fmt"
)

type customErr struct {
  info string
}

// if we want to attach a builtin,error interface to customErr
// we need to attach method signature "Error() string" to the struct
// customErr.
// by passing customErr as receiver to a func we have done the above mention
func (ce customErr) Error() string {
  return fmt.Sprintf("custom erro: %v", ce.info)
}

func main() {
	c1 := customErr{
    info : "this is the custom error",
  }
	foo(c1)
}

func foo(e error) {
  fmt.Println("foo ran -", e)

  // assertion
  // we do fmt.Println(e,info) it won't work as the compiler see
  // info as the field of customErr not error
  // so to get the info field we have to assert that e.(customErr)
  // is a type of customErr and it has a field info
  //fmt.Println("foo ran ", e, "\n", e.(customErr).info)
}
