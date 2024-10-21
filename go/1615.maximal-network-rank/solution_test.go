package maximalnetworkrank

import "testing"

func Test_maximalNetworkRank(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				n: 4,
				roads: [][]int{
					{0, 1}, {0, 3}, {1, 2}, {1, 3},
				},
			},
			want: 4,
		},
		{
			name: "tc2",
			args: args{
				n: 5,
				roads: [][]int{
					{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4},
				},
			},
			want: 5,
		},
		{
			name: "tc3",
			args: args{
				n: 8,
				roads: [][]int{
					{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalNetworkRank(tt.args.n, tt.args.roads); got != tt.want {
				t.Errorf("maximalNetworkRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
