package leetcodesolution

import "testing"

func Test_subsetXORSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 3},
			},
			wantAns: 6,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{5, 1, 6},
			},
			wantAns: 28,
		},
		{
			name: "tc1",
			args: args{
				nums: []int{3, 4, 5, 6, 7, 8},
			},
			wantAns: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subsetXORSum(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("subsetXORSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
