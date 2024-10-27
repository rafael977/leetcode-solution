// Created by Rafael Shen at 2024/10/27 10:55
// leetgo: 1.4.9
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/

package countsquaresubmatriceswithallones

import "testing"

func Test_countSquares(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				matrix: [][]int{
					{0, 1, 1, 1},
					{1, 1, 1, 1},
					{0, 1, 1, 1},
				},
			},
			wantAns: 15,
		},
		{
			name: "tc2",
			args: args{
				matrix: [][]int{
					{1, 0, 1},
					{1, 1, 0},
					{1, 1, 0},
				},
			},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSquares(tt.args.matrix); gotAns != tt.wantAns {
				t.Errorf("countSquares() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
