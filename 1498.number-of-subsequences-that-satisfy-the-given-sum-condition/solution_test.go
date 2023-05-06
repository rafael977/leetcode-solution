package numberofsubsequencesthatsatisfythegivensumcondition

import "testing"

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				nums:   []int{3, 5, 6, 7},
				target: 9,
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				nums:   []int{3, 3, 6, 8},
				target: 10,
			},
			wantAns: 6,
		},
		{
			name: "tc3",
			args: args{
				nums:   []int{2, 3, 3, 4, 6, 7},
				target: 12,
			},
			wantAns: 61,
		},
		{
			name: "tc4",
			args: args{
				nums:   []int{5, 2, 4, 1, 7, 6, 8},
				target: 16,
			},
			wantAns: 127,
		},
		{
			name: "tc5",
			args: args{
				nums:   []int{11, 9, 11, 13, 4, 13, 9, 9, 10, 1, 9, 6, 5, 2, 1, 8, 5, 5},
				target: 22,
			},
			wantAns: 262119,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSubseq(tt.args.nums, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("numSubseq() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
