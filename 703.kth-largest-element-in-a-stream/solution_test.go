package kthlargestelementinastream

import (
	"fmt"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	kl := Constructor(3, []int{4, 5, 8, 2})

	tests := []struct {
		val  int
		want int
	}{
		{
			val:  3,
			want: 4,
		},
		{
			val:  5,
			want: 5,
		},
		{
			val:  10,
			want: 5,
		},
		{
			val:  9,
			want: 8,
		},
		{
			val:  4,
			want: 8,
		},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("tc%d", i), func(t *testing.T) {
			if v := kl.Add(tc.val); v != tc.want {
				t.Errorf("add return %d, want %d", v, tc.want)
			}
		})
	}
}
