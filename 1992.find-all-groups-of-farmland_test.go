package leetcodesolution

import (
	"reflect"
	"testing"
)

func Test_findFarmland(t *testing.T) {
	type args struct {
		land [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				[][]int{
					{1, 0, 0},
					{0, 1, 1},
					{0, 1, 1},
				},
			},
			wantAns: [][]int{
				{0, 0, 0, 0},
				{1, 1, 2, 2},
			},
		},
		{
			name: "tc2",
			args: args{
				[][]int{
					{1, 1},
					{1, 1},
				},
			},
			wantAns: [][]int{
				{0, 0, 1, 1},
			},
		},
		{
			name: "tc3",
			args: args{
				[][]int{
					{0},
				},
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findFarmland(tt.args.land); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findFarmland() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
