package combinationsumivq

import "testing"

func Test_combinationSum4(t *testing.T) {
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
				nums:   []int{1, 2, 3},
				target: 4,
			},
			wantAns: 7,
		},
		{
			name: "tc2",
			args: args{
				nums:   []int{9},
				target: 3,
			},
			wantAns: 0,
		},
		{
			name: "tc3",
			args: args{
				nums:   []int{2, 1, 3},
				target: 35,
			},
			wantAns: 1132436852,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum4(tt.args.nums, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("combinationSum4() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
