package wigglesubsequence

import "testing"

func Test_wiggleMaxLength(t *testing.T) {
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
				nums: []int{1, 7, 4, 9, 2, 5},
			},
			wantAns: 6,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			},
			wantAns: 7,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := wiggleMaxLength(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("wiggleMaxLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
