package leetcodesolution

import "testing"

func Test_minDifference(t *testing.T) {
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
				nums: []int{5, 3, 2, 4},
			},
			wantAns: 0,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 5, 0, 10, 14},
			},
			wantAns: 1,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{3, 100, 20},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minDifference(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
