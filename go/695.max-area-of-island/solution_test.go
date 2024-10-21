package maxareaofisland

import "testing"

func Test_maxAreaOfIsland(t *testing.T) {
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
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
				},
			},
			wantAns: 6,
		},
		{
			name: "tc2",
			args: args{
				grid: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			wantAns: 0,
		},
		{
			name: "tc3",
			args: args{
				grid: [][]int{
					{1, 1},
					{1, 0},
				},
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxAreaOfIsland(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("maxAreaOfIsland() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
