package reducingdishes

import "testing"

func Test_maxSatisfaction(t *testing.T) {
	type args struct {
		satisfaction []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				satisfaction: []int{-1, -8, 0, 5, -9},
			},
			wantAns: 14,
		},
		{
			name: "tc2",
			args: args{
				satisfaction: []int{4, 3, 2},
			},
			wantAns: 20,
		},
		{
			name: "tc3",
			args: args{
				satisfaction: []int{-1, -4, -5},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxSatisfaction(tt.args.satisfaction); gotAns != tt.wantAns {
				t.Errorf("maxSatisfaction() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
