package main

import "testing"

func Test_addDigits(t *testing.T) {
	tcs := map[string]struct {
		num      int
		expected int
	}{
		"test case 1": {
			num:      38,
			expected: 2,
		},
		"test case 2": {
			num:      0,
			expected: 0,
		},
		"test case 3": {
			num:      1234567,
			expected: 1,
		},
		"test case 4": {
			num:      123456789,
			expected: 9,
		},
	}

	for n, tc := range tcs {
		tc := tc
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			r := addDigits(tc.num)
			if r != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, r)
			}
		})
	}
}
