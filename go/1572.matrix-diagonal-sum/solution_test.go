package matrixdiagonalsum

import "testing"

func Test_diagonalSum(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				mat: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantAns: 25,
		},
		{
			name: "tc2",
			args: args{
				mat: [][]int{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
				},
			},
			wantAns: 8,
		},
		{
			name: "tc1",
			args: args{
				mat: [][]int{
					{5},
				},
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := diagonalSum(tt.args.mat); gotAns != tt.wantAns {
				t.Errorf("diagonalSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
