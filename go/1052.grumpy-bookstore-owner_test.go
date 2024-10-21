package leetcodesolution

import "testing"

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
				grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
				minutes:   3,
			},
			wantAns: 16,
		},
		{
			name: "tc2",
			args: args{
				customers: []int{1},
				grumpy:    []int{0},
				minutes:   1,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes); gotAns != tt.wantAns {
				t.Errorf("maxSatisfied() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
