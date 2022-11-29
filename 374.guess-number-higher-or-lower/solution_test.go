package guessnumberhigherorlower

import (
	"testing"
)

func Test_guessNumber(t *testing.T) {
	type args struct {
		n    int
		pick int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				n:    10,
				pick: 6,
			},
			want: 6,
		},
		{
			name: "tc2",
			args: args{
				n:    1,
				pick: 1,
			},
			want: 1,
		},
		{
			name: "tc3",
			args: args{
				n:    2,
				pick: 1,
			},
			want: 1,
		},
		{
			name: "tc4",
			args: args{
				n:    2,
				pick: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guess = getGuess(tt.args.pick)
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getGuess(pick int) func(num int) int {
	return func(num int) int {
		if pick == num {
			return 0
		} else if pick > num {
			return 1
		} else {
			return -1
		}
	}
}
