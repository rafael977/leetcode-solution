package jumpgamevi

import "testing"

func Test_maxResult(t *testing.T) {
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
				nums: []int{1, -1, -2, 4, -7, 3},
				k:    2,
			},
			wantAns: 7,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{10, -5, -2, 4, 0, 3},
				k:    3,
			},
			wantAns: 17,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{1, -5, -20, 4, -1, 3, -6, -3},
				k:    2,
			},
			wantAns: 0,
		},
		{
			name: "tc4",
			args: args{
				nums: []int{100, -1, -100, -1, 100},
				k:    2,
			},
			wantAns: 198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxResult(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxResult() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
