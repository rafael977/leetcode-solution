package leetcodesolution

import "testing"

func Test_pivotInteger(t *testing.T) {
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
			args: args{n: 8},
			want: 6,
		},
		{
			name: "tc2",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "tc3",
			args: args{n: 4},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotInteger(tt.args.n); got != tt.want {
				t.Errorf("pivotInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
