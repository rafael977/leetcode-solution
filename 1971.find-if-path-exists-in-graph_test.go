package leetcodesolution

import "testing"

func Test_validPath(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1}, {1, 2}, {2, 0},
				},
				source:      0,
				destination: 2,
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				n: 6,
				edges: [][]int{
					{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3},
				},
				source:      0,
				destination: 5,
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				n: 10,
				edges: [][]int{
					{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5},
				},
				source:      7,
				destination: 5,
			},
			want: true,
		},
		{
			name: "tc4",
			args: args{
				n:           1,
				edges:       [][]int{},
				source:      0,
				destination: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPath(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("validPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
