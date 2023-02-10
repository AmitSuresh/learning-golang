package main

import "testing"

func TestMySum(t *testing.T) {
	want := 100
	if got := mySum(50, 50); got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}
