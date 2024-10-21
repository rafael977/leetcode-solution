package maximumswap

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{num: 2736},
			want: 7236,
		},
		{
			name: "tc2",
			args: args{num: 9973},
			want: 9973,
		},
		{
			name: "tc3",
			args: args{num: 98368},
			want: 98863,
		},
		{
			name: "tc4",
			args: args{num: 1993},
			want: 9913,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
