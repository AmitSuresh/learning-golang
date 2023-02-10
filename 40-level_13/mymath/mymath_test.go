package mymath

import (
	"fmt"
	"testing"
)

var slice = []int{70, 70, 50, 20, 10, 10, 60, 70, 20, 40, 50, 80, 10, 50, 90, 100}

func TestCenteredAvg(t *testing.T) {
	got := CenteredAvg(slice)
	if got != 49.285714285714285 {
		t.Errorf("got: %v, want: %v", got, 70)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg(slice))
	//Output:
	//49.285714285714285
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg(slice)
	}
}

func TestCenteredAvg_first(t *testing.T) {
	type testdataanswer struct {
		data   []int
		answer float64
	}
	testslices := []testdataanswer{
		{[]int{70, 70, 50, 20, 10, 10, 60, 70, 20, 40, 50, 80, 10, 50, 90, 100}, 49.285714285714285},
		{[]int{10, 20, 40, 60, 80}, 40},
		{[]int{1, 2, 4, 6, 8, 10}, 5},
		{[]int{1, 3, 6, 9, 12, 15}, 7.5},
	}
	for _, slice := range testslices {
		got := CenteredAvg(slice.data)
		if got != slice.answer {
			t.Errorf("got: %v, want: %v", got, slice.answer)
		}
	}
}
