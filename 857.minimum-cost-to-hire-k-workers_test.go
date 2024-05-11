package leetcodesolution

import (
	"math"
	"testing"
)

func Test_mincostToHireWorkers(t *testing.T) {
	type args struct {
		quality []int
		wage    []int
		k       int
	}
	tests := []struct {
		name    string
		args    args
		wantAns float64
	}{
		{
			name: "tc1",
			args: args{
				quality: []int{10, 20, 5},
				wage:    []int{70, 50, 30},
				k:       2,
			},
			wantAns: 105,
		},
		{
			name: "tc2",
			args: args{
				quality: []int{3, 1, 10, 10, 1},
				wage:    []int{4, 8, 2, 2, 7},
				k:       3,
			},
			wantAns: 30.66667,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := mincostToHireWorkers(tt.args.quality, tt.args.wage, tt.args.k); math.Abs(gotAns-tt.wantAns) > 1e-5 {
				t.Errorf("mincostToHireWorkers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
