package nondecreasingarray

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{4, 2, 3},
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{4, 2, 1},
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{5, 7, 1, 8},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
