package leetcodesolution

import "testing"

func Test_findTheCity(t *testing.T) {
	type args struct {
		n                 int
		edges             [][]int
		distanceThreshold int
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
				edges: [][]int{
					{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1},
				},
				distanceThreshold: 4,
			},
			want: 3,
		},
		{
			name: "tc2",
			args: args{
				n: 5,
				edges: [][]int{
					{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1},
				},
				distanceThreshold: 2,
			},
			want: 0,
		},
		{
			name: "tc3",
			args: args{
				n: 6,
				edges: [][]int{
					{0, 1, 10}, {0, 2, 1}, {2, 3, 1}, {1, 3, 1}, {1, 4, 1}, {4, 5, 10},
				},
				distanceThreshold: 20,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheCity(tt.args.n, tt.args.edges, tt.args.distanceThreshold); got != tt.want {
				t.Errorf("findTheCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
