package reorderedpowerof2

import "testing"

func Test_reorderedPowerOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{n: 1},
			want: true,
		},
		{
			name: "tc2",
			args: args{n: 10},
			want: false,
		},
		{
			name: "tc3",
			args: args{n: 4201},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
