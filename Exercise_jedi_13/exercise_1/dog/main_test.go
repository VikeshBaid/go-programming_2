package dog

import (
  "fmt"
  "testing"
)

func TestYears(t *testing.T) {
  x := Years(1)
  if x != 7 {
    t.Error("Got", x, "Expected", 7)
  }
}

func TestYearstwo(t *testing.T) {
  x := YearsTwo(1)
  if x != 7 {
    t.Error("Got", x, "Expected", 7)
  }
}

func ExampleYears()  {
  fmt.Println(Years(2))
  // Output:
  // 14
}

func ExampleYearsTwo()  {
  fmt.Println(YearsTwo(2))
  // Output:
  // 14
}

func BenchmarkYears(b *testing.B)  {
  for i := 0; i < b.N; i++ {
    Years(1)
  }
}

func BenchmarkYearstwo(b *testing.B)  {
  for i := 0; i < b.N; i++ {
    YearsTwo(1)
  }
}
