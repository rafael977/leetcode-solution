package countnegativenumbersinasortedmatrix

import "testing"

func Test_countNegatives(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				grid: [][]int{
					{4, 3, 2, -1},
					{3, 2, 1, -1},
					{1, 1, -1, -2},
					{-1, -1, -2, -3},
				},
			},
			wantAns: 8,
		},
		{
			name: "tc2",
			args: args{
				grid: [][]int{
					{3, 2},
					{1, 0},
				},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countNegatives(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("countNegatives() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
