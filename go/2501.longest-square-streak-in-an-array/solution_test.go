// Created by Rafael Shen at 2024/10/28 10:10
// leetgo: 1.4.9
// https://leetcode.com/problems/longest-square-streak-in-an-array/

package longestsquarestreakinanarray

import "testing"

func Test_longestSquareStreak(t *testing.T) {
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
				nums: []int{4, 3, 6, 16, 8, 2},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{2, 3, 5, 6, 7},
			},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSquareStreak(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("longestSquareStreak() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
