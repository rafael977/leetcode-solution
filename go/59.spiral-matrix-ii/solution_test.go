package spiralmatrixii

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				n: 3,
			},
			wantAns: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "tc2",
			args: args{
				n: 1,
			},
			wantAns: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := generateMatrix(tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("generateMatrix() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
