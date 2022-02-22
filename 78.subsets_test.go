package main

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 2, 3},
			},
			wantAns: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{0},
			},
			wantAns: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subsets(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("subsets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
