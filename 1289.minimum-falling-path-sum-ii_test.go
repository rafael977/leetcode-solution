package leetcodesolution

import "testing"

func Test_minFallingPathSum(t *testing.T) {
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
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "tc2",
			args: args{
				grid: [][]int{
					{7},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
