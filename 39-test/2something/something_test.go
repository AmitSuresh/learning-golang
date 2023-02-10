package something

import (
	"testing"
)

func TestSomething(t *testing.T) {
	test_slice := []int{10, 90}

	got := Sum(test_slice...)
	want := 100
	if got != want {
		t.Error("got: ", got, "expected: ", want)
	}
}
