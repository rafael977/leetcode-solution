package main

import "testing"

func Test_brokenCalc(t *testing.T) {
	type args struct {
		startValue int
		target     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				startValue: 2,
				target:     3,
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				startValue: 5,
				target:     8,
			},
			want: 2,
		},
		{
			name: "tc3",
			args: args{
				startValue: 3,
				target:     10,
			},
			want: 3,
		},
		{
			name: "tc4",
			args: args{
				startValue: 1024,
				target:     1,
			},
			want: 1023,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brokenCalc(tt.args.startValue, tt.args.target); got != tt.want {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
