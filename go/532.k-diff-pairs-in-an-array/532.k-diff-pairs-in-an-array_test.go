package main

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{3, 1, 4, 1, 5},
				k:    2,
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			wantAns: 4,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{3, 1, 4, 1, 5},
				k:    0,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findPairs(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("findPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
