package shortestpathinbinarymatrix

import "testing"

func Test_shortestPathBinaryMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				grid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				grid: [][]int{
					{0, 0, 0},
					{1, 1, 0},
					{1, 1, 0},
				},
			},
			want: 4,
		},
		{
			name: "tc3",
			args: args{
				grid: [][]int{
					{1, 0, 0},
					{1, 1, 0},
					{1, 1, 0},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathBinaryMatrix(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
