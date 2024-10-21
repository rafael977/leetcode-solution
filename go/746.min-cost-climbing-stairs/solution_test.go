package mincostclimbingstairs

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				cost: []int{10, 15, 20},
			},
			wantAns: 15,
		},
		{
			name: "tc2",
			args: args{
				cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minCostClimbingStairs(tt.args.cost); gotAns != tt.wantAns {
				t.Errorf("minCostClimbingStairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
