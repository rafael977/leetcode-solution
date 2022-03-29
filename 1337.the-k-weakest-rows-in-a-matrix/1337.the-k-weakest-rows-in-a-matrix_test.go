package main

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				mat: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 0},
					{1, 0, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{1, 1, 1, 1, 1},
				},
				k: 3,
			},
			wantAns: []int{2, 0, 3},
		},
		{
			name: "tc2",
			args: args{
				mat: [][]int{
					{1, 0, 0, 0},
					{1, 1, 1, 1},
					{1, 0, 0, 0},
					{1, 0, 0, 0},
				},
				k: 2,
			},
			wantAns: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("kWeakestRows() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
