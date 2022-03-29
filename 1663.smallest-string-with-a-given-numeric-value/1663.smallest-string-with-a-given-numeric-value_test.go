package main

import "testing"

func Test_getSmallestString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				n: 3,
				k: 27,
			},
			want: "aay",
		},
		{
			name: "tc2",
			args: args{
				n: 5,
				k: 73,
			},
			want: "aaszz",
		},
		{
			name: "tc3",
			args: args{
				n: 24,
				k: 552,
			},
			want: "aadzzzzzzzzzzzzzzzzzzzzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
