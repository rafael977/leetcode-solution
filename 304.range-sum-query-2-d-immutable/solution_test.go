package rangesumquery2dimmutable

import "testing"

func TestNumMatrix_SumRegion(t *testing.T) {
	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name    string
		obj     NumMatrix
		args    args
		wantSum int
	}{
		{
			name: "tc1",
			obj: Constructor([][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}),
			args: args{
				row1: 2,
				col1: 1,
				row2: 4,
				col2: 3,
			},
			wantSum: 8,
		},
		{
			name: "tc2",
			obj: Constructor([][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}),
			args: args{
				row1: 1,
				col1: 1,
				row2: 2,
				col2: 2,
			},
			wantSum: 11,
		},
		{
			name: "tc3",
			obj: Constructor([][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}),
			args: args{
				row1: 1,
				col1: 2,
				row2: 2,
				col2: 4,
			},
			wantSum: 12,
		},
		{
			name: "tc4",
			obj: Constructor([][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			}),
			args: args{
				row1: 0,
				col1: 1,
				row2: 2,
				col2: 2,
			},
			wantSum: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := tt.obj.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); gotSum != tt.wantSum {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
