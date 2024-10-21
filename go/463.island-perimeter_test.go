package leetcodesolution

import "testing"

func Test_islandPerimeter(t *testing.T) {
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
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{0, 1, 0, 0},
					{1, 1, 0, 0},
				},
			},
			wantAns: 16,
		},
		{
			name: "tc2",
			args: args{
				grid: [][]int{
					{1},
				},
			},
			wantAns: 4,
		},
		{
			name: "tc3",
			args: args{
				grid: [][]int{
					{1, 0},
				},
			},
			wantAns: 4,
		},
		{
			name: "tc4",
			args: args{
				grid: [][]int{
					{0},
					{1},
				},
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := islandPerimeter(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("islandPerimeter() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
