package besttimetobuyandsellstock

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxProfit(tt.args.prices); gotAns != tt.wantAns {
				t.Errorf("maxProfit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
