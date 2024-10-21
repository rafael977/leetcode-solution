package minimumaveragedifference

import "testing"

func Test_minimumAverageDifference(t *testing.T) {
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
				nums: []int{2, 5, 3, 9, 5, 3},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{0},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumAverageDifference(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minimumAverageDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
