package leetcodesolution

import "testing"

func Test_findMaxK(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{-1, 2, -3, 3},
			},
			want: 3,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{-1, 10, 6, 7, -7, 1},
			},
			want: 7,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{-10, 8, 6, 7, -2, -3},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxK(tt.args.nums); got != tt.want {
				t.Errorf("findMaxK() = %v, want %v", got, tt.want)
			}
		})
	}
}
