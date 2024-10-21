package transposematrix

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name   string
		args   args
		wantTm [][]int
	}{
		{
			name: "tc1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantTm: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			name: "tc2",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			wantTm: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTm := transpose(tt.args.matrix); !reflect.DeepEqual(gotTm, tt.wantTm) {
				t.Errorf("transpose() = %v, want %v", gotTm, tt.wantTm)
			}
		})
	}
}
