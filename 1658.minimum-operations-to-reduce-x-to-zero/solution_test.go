package minimumoperationstoreducextozero

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 1, 4, 2, 3},
				x:    5,
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{5, 6, 7, 8, 9},
				x:    4,
			},
			want: -1,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{3, 2, 20, 1, 1, 3},
				x:    10,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
