package main

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
		{
			name: "tc2",
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
		{
			name: "tc3",
			args: args{
				num: "10",
				k:   2,
			},
			want: "0",
		},
		{
			name: "tc4",
			args: args{
				num: "9",
				k:   1,
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
