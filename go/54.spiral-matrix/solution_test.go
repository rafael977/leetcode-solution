package spiralmatrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
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
			wantAns: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "tc2",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			},
			wantAns: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := spiralOrder(tt.args.matrix); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("spiralOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
