package word

import (
	"fmt"
	"testing"

	"github.com/AmitSuresh/learning-golang/40-level_13/second/quote"
)

func TestUseCount(t *testing.T) {
	got := UseCount(quote.SunAlso)
	if got["one"] != 8 {
		t.Errorf("got: %v, want: %v", got, 8)
	}
}

func ExampleUseCount() {
	newmap := UseCount(quote.SunAlso)
	fmt.Println(newmap["one"])
	//Output:
	//8
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func TestCount(t *testing.T) {
	got := Count(quote.SunAlso)
	if got != 1349 {
		t.Errorf("got: %v, want: %v", got, 70)
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	//Output:
	//1349
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func TestUseCount_first(t *testing.T) {
	strInt := UseCount("one two two three three three")
	for strs, ints := range strInt {
		switch strs {
		case "one":
			if ints != 1 {
				t.Errorf("got: %v, want: %v", ints, 1)
			}
		case "two":
			if ints != 2 {
				t.Errorf("got: %v, want: %v", ints, 1)
			}
		case "three":
			if ints != 3 {
				t.Errorf("got: %v, want: %v", ints, 1)
			}
		}
	}
}
