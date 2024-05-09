package leetcodesolution

import "testing"

func Test_maximumHappinessSum(t *testing.T) {
	type args struct {
		happiness []int
		k         int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "tc1",
			args: args{
				happiness: []int{1, 2, 3},
				k:         2,
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				happiness: []int{1, 1, 1, 1},
				k:         2,
			},
			wantAns: 1,
		},
		{
			name: "tc3",
			args: args{
				happiness: []int{2, 3, 4, 5},
				k:         1,
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumHappinessSum(tt.args.happiness, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maximumHappinessSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
