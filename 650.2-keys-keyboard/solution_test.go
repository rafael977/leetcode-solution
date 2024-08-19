package keyskeyboard

import "testing"

func Test_minSteps(t *testing.T) {
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
			want: 3,
		},
		{
			name: "tc2",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "tc3",
			args: args{
				n: 6,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
