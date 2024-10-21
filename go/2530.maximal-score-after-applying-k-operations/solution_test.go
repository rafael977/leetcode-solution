package maximalscoreafterapplyingkoperations

import "testing"

func Test_maxKelements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{10, 10, 10, 10, 10, 10},
				k:    5,
			},
			want: 50,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 10, 3, 3, 3},
				k:    3,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKelements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxKelements() = %v, want %v", got, tt.want)
			}
		})
	}
}
