package singlethreadedcpu

import (
	"reflect"
	"testing"
)

func Test_getOrder(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				tasks: [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}},
			},
			wantAns: []int{0, 2, 3, 1},
		},
		{
			name: "tc2",
			args: args{
				tasks: [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}},
			},
			wantAns: []int{4, 3, 2, 0, 1},
		},
		{
			name: "tc3",
			args: args{
				tasks: [][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}},
			},
			wantAns: []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7},
		},
		{
			name: "tc4",
			args: args{
				tasks: [][]int{{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 11}},
			},
			wantAns: []int{5, 0, 1, 3, 2, 4},
		},
		{
			name: "tc5",
			args: args{
				tasks: [][]int{{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 1}},
			},
			wantAns: []int{5, 0, 1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getOrder(tt.args.tasks); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
