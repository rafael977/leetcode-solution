package maxnumberofksumpairs

import "testing"

func Test_maxOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    5,
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{3, 1, 3, 4, 3},
				k:    6,
			},
			wantAns: 1,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2},
				k:    3,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxOperations(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
