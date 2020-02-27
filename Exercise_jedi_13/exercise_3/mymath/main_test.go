package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

  type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{2, 3, 4, 5, 6, 7}, 4.5},
		test{[]int{1, 4, 6, 8, 100}, 6},
		test{[]int{0, 8, 10, 1000}, 9},
		test{[]int{9000, 4, 10, 8, 6, 12}, 9},
	}

	for _, v := range tests {
		ca := CenteredAvg(v.data)
		if ca != v.answer {
			t.Error("got", ca, "expected", v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{2, 3, 4, 5, 6, 7}))
	// Output:
	// 4.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{2, 3, 4, 5, 6, 7})
	}
}
