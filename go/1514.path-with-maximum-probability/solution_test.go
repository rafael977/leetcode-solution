package pathwithmaximumprobability

import (
	"math"
	"testing"
)

func Test_maxProbability(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		succProb []float64
		start    int
		end      int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "tc1",
			args: args{
				n:        3,
				edges:    [][]int{{0, 1}, {1, 2}, {0, 2}},
				succProb: []float64{0.5, 0.5, 0.2},
				start:    0,
				end:      2,
			},
			want: 0.25,
		},
		{
			name: "tc2",
			args: args{
				n:        3,
				edges:    [][]int{{0, 1}, {1, 2}, {0, 2}},
				succProb: []float64{0.5, 0.5, 0.3},
				start:    0,
				end:      2,
			},
			want: 0.3,
		},
		{
			name: "tc2",
			args: args{
				n:        3,
				edges:    [][]int{{0, 1}},
				succProb: []float64{0.5},
				start:    0,
				end:      2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProbability(tt.args.n, tt.args.edges, tt.args.succProb, tt.args.start, tt.args.end); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("maxProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
