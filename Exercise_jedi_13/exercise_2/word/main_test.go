package word

import (
  "fmt"
	"testing"
  "github.com/go-programming_2/Exercise_jedi_13/exercise_2/quote"
)

func TestCount(t *testing.T)  {
  n := Count("this is a string")
  if n != 4 {
    t.Error("got", n, "Expected", 4)
  }
}

func TestUseCount(t *testing.T)  {
  m := UseCount("one two two three three three")
  for k, v := range m {
    switch k {
    case "one" :
      if v != 1 {
        t.Error("Got", v, "Expexted", 1)
      }
    case "two" :
      if v != 2 {
        t.Error("Got", v, "Expexted", 2)
      }
    case "three" :
      if v != 3 {
        t.Error("Got", v, "Expexted", 3)
      }
    }
  }
}

func ExampleCount()  {
  fmt.Println(Count("This is a string"))
  // Output
  // 4
}

func ExampleUseCount()  {
  fmt.Println(UseCount("this this is a string string string"))
  // Output
  // 2 This
  // 1 is
  // 1 a
  // 3 string
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
