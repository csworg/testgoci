package main

import "testing"

func TestSum(t *testing.T) {
	sum := sum(2, 2)
	result := 4
	if sum != result {
		t.Fatalf("ERROR => sum does not equal desired result")
	}
}
