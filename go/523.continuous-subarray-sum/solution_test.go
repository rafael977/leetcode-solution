package continuoussubarraysum

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{23, 2, 4, 6, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    13,
			},
			want: false,
		},
		{
			name: "tc4",
			args: args{
				nums: []int{23, 2, 4, 6, 6},
				k:    7,
			},
			want: true,
		},
		{
			name: "tc5",
			args: args{
				nums: []int{1, 0},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
