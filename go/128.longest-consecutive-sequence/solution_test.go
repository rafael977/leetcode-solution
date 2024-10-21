package longestconsecutivesequence

import "testing"

func Test_longestConsecutive(t *testing.T) {
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
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			wantAns: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestConsecutive(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("longestConsecutive() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
