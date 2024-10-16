package longesthappystring

import "testing"

func Test_longestDiverseString(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				a: 1,
				b: 1,
				c: 7,
			},
			want: "ccaccbcc",
		},
		{
			name: "tc2",
			args: args{
				a: 7,
				b: 1,
				c: 0,
			},
			want: "aabaa",
		},
		{
			name: "tc3",
			args: args{
				a: 2,
				b: 2,
				c: 1,
			},
			want: "aabbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDiverseString(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("longestDiverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
