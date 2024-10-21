package leetcodesolution

import (
	"slices"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			wantAns: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name: "tc2",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			wantAns: [][]int{
				{1, 2, 2},
				{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum2(tt.args.candidates, tt.args.target); !slices.EqualFunc(gotAns, tt.wantAns, func(e1, e2 []int) bool {
				if len(gotAns) != len(tt.wantAns) ||
					!slices.ContainsFunc(tt.wantAns, func(e []int) bool { return slices.Equal(e, e1) }) ||
					!slices.ContainsFunc(gotAns, func(e []int) bool { return slices.Equal(e, e2) }) {
					return false
				}
				return true
			}) {
				t.Errorf("combinationSum2() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
