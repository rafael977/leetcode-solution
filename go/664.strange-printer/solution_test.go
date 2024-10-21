package strangeprinter

import "testing"

func Test_strangePrinter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				s: "aaabbb",
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				s: "aba",
			},
			want: 2,
		},
		{
			name: "tc3",
			args: args{
				s: "tbgtgb",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.args.s); got != tt.want {
				t.Errorf("strangePrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}
