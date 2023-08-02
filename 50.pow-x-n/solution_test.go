package powxn

import (
	"math"
	"testing"
)

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "tc1",
			args: args{
				x: 2.0,
				n: 10,
			},
			want: 1024,
		},
		{
			name: "tc2",
			args: args{
				x: 2.1,
				n: 3,
			},
			want: 9.261,
		},
		{
			name: "tc3",
			args: args{
				x: 2.0,
				n: -2,
			},
			want: 0.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); math.Abs(got-tt.want) > 10e-9 {
				t.Errorf("myPow() = %v, want %v, diff %v", got, tt.want, math.Abs(got-tt.want))
			}
		})
	}
}
