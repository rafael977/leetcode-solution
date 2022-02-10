package main

import "testing"

func Test_subarraySum(t *testing.T) {
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			wantAns: 2,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{1, -1, 0},
				k:    0,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subarraySum(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("subarraySum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
