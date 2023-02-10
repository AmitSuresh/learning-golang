package main

import "testing"

func TestMySums(t *testing.T) {
	type test1 struct {
		data   []int
		answer int
	}

	tests := []test1{
		{[]int{1, 1}, 2},
		{[]int{5, 5}, 10},
		{[]int{10, 10}, 20},
	}
	for _, v := range tests {
		got := mySum(v.data...)
		if got != v.answer {
			t.Errorf("Got: %v, Want: %v", got, v.answer)
		}
	}
}
