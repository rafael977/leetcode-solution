package main

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "tc1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "tc2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "tc3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !sliceEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sliceEqual(g, w [][]int) bool {
	for _, gi := range g {
		found := false
		for _, wi := range w {
			if reflect.DeepEqual(gi, wi) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
