package leetcodesolution

import "testing"

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				[]int{2, 3, 1, 6, 7},
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				[]int{1, 1, 1, 1, 1},
			},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countTriplets(tt.args.arr); gotAns != tt.wantAns {
				t.Errorf("countTriplets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
