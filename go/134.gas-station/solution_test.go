package gasstation

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
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
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := canCompleteCircuit(tt.args.gas, tt.args.cost); gotAns != tt.wantAns {
				t.Errorf("canCompleteCircuit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
