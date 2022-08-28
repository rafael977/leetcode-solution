package splitarrayintoconsecutivesubsequences

import "testing"

func Test_isPossible(t *testing.T) {
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
				[]int{1, 2, 3, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				[]int{1, 2, 3, 3, 4, 4, 5, 5},
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				[]int{1, 2, 3, 4, 4, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.nums); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
