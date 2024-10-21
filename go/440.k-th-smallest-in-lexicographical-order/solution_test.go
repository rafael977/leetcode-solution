package kthsmallestinlexicographicalorder

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				n: 13,
				k: 2,
			},
			want: 10,
		},
		{
			name: "tc2",
			args: args{
				n: 1,
				k: 1,
			},
			want: 1,
		},
		{
			name: "tc3",
			args: args{
				n: 10000,
				k: 10000,
			},
			want: 9999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
