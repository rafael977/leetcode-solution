// Created by Rafael Shen at 2024/10/30 10:17
// leetgo: 1.4.9
// https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/

package minimumnumberofremovalstomakemountainarray

import "testing"

func Test_minimumMountainRemovals(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 3, 1},
			},
			wantAns: 0,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{2, 1, 1, 5, 6, 2, 3, 1},
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumMountainRemovals(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minimumMountainRemovals() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
