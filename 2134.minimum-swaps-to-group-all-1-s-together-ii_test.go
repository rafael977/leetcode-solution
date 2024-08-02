package leetcodesolution

import "testing"

func Test_minSwaps(t *testing.T) {
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
				nums: []int{0, 1, 0, 1, 1, 0, 0},
			},
			wantAns: 1,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{0, 1, 1, 1, 0, 0, 1, 1, 0},
			},
			wantAns: 2,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{1, 1, 0, 0, 1},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minSwaps(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minSwaps() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
