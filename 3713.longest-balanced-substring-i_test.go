package main

import "testing"

func Test_longestBalanced(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{name: "tc1", s: "abbac", want: 4},
		{name: "tc2", s: "zzabccy", want: 4},
		{name: "tc3", s: "aba", want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestBalanced(tt.s)
			if got != tt.want {
				t.Errorf("longestBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
