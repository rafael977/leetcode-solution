package leetcodesolution

import "testing"

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				tickets: []int{2, 3, 2},
				k:       2,
			},
			want: 6,
		},
		{
			name: "tc2",
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       0,
			},
			want: 8,
		},
		{
			name: "tc3",
			args: args{
				tickets: []int{
					15, 66, 3, 47, 71, 27, 54, 43, 97, 34, 94, 33, 54,
					26, 15, 52, 20, 71, 88, 42, 50, 6, 66, 88, 36, 99,
					27, 82, 7, 72,
				},
				k: 18,
			},
			want: 1457,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
