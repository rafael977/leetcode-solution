package main

import "testing"

func Test_twoCitySchedCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				costs: [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
			},
			wantAns: 110,
		},
		{
			name: "tc2",
			args: args{
				costs: [][]int{
					{259, 770},
					{448, 54},
					{926, 667},
					{184, 139},
					{840, 118},
					{577, 469},
				},
			},
			wantAns: 1859,
		},
		{
			name: "tc3",
			args: args{
				costs: [][]int{
					{515, 563},
					{451, 713},
					{537, 709},
					{343, 819},
					{855, 779},
					{457, 60},
					{650, 359},
					{631, 42},
				},
			},
			wantAns: 3086,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := twoCitySchedCost(tt.args.costs); gotAns != tt.wantAns {
				t.Errorf("twoCitySchedCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
