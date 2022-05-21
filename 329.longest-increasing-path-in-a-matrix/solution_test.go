package longestincreasingpathinamatrix

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				matrix: [][]int{
					{9, 9, 4},
					{6, 6, 8},
					{2, 1, 1},
				},
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				matrix: [][]int{
					{3, 4, 5},
					{3, 2, 6},
					{2, 2, 1},
				},
			},
			wantAns: 4,
		},
		{
			name: "tc3",
			args: args{
				matrix: [][]int{
					{1},
				},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestIncreasingPath(tt.args.matrix); gotAns != tt.wantAns {
				t.Errorf("longestIncreasingPath() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
