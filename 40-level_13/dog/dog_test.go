package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	got := Years(10)
	if got != 70 {
		t.Errorf("got: %v, want: %v", got, 70)
	}
}

const human_years = 10

func ExampleYears() {
	fmt.Println(Years(human_years))
	//Output:
	//70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(human_years)
	}
}
func TestYearsTwo(t *testing.T) {
	got := Years(10)
	if got != 70 {
		t.Errorf("got: %v, want: %v", got, 70)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(human_years))
	//Output:
	//70
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(human_years)
	}
}
