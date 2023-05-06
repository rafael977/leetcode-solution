package bulbswitcher

import "testing"

func Test_bulbSwitch(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				n: 3,
			},
			want: 1,
		},
		{
			name: "tc2",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "tc3",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "tc4",
			args: args{
				n: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulbSwitch(tt.args.n); got != tt.want {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
