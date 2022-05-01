package mincosttoconnectallpoints

import "testing"

func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			},
			wantAns: 20,
		},
		{
			name: "tc2",
			args: args{
				points: [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			},
			wantAns: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minCostConnectPoints(tt.args.points); gotAns != tt.wantAns {
				t.Errorf("minCostConnectPoints() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
