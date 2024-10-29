// Created by Rafael Shen at 2024/10/29 21:12
// leetgo: 1.4.9
// https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/

package maximumnumberofmovesinagrid

import "testing"

func Test_maxMoves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				grid: [][]int{
					{2, 4, 3, 5},
					{5, 4, 9, 3},
					{3, 4, 2, 11},
					{10, 9, 13, 15},
				},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				grid: [][]int{
					{3, 2, 4},
					{2, 1, 9},
					{1, 1, 7},
				},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxMoves(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("maxMoves() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
