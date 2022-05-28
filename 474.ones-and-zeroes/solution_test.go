package onesandzeroes

import "testing"

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    5,
				n:    3,
			},
			want: 4,
		},
		{
			name: "tc2",
			args: args{
				strs: []string{"10", "0", "1"},
				m:    1,
				n:    1,
			},
			want: 2,
		},
		{
			name: "tc3",
			args: args{
				strs: []string{"0", "11", "1000", "01", "0", "101", "1", "1", "1", "0", "0", "0", "0", "1", "0", "0110101", "0", "11", "01", "00", "01111", "0011", "1", "1000", "0", "11101", "1", "0", "10", "0111"},
				m:    9,
				n:    80,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
