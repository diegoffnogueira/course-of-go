package main

import (
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestMySum(t *testing.T) {

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, value := range tests {
		x := mySum(value.data...)
		if x != value.answer {
			t.Error("Expected", value.answer, "Got", x)
		}
	}
}
