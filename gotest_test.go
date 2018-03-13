package gotest

import "testing"

func TestAdd(t *testing.T) {
	var tests =  []struct {
	a, b int
	want int
	} {
		{2, 3, 5},
		{0, 0, 0},
	}

	for _, test := range tests {
		if got := add(test.a, test.b); got != test.want {
			t.Errorf("add(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}
}
