package uniquepathsii

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				obstacleGrid: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				obstacleGrid: [][]int{
					{0, 1},
					{0, 0},
				},
			},
			want: 1,
		},
		{
			name: "tc3",
			args: args{
				obstacleGrid: [][]int{
					{0, 1, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
