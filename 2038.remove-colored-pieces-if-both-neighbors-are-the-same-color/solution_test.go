package removecoloredpiecesifbothneighborsarethesamecolor

import "testing"

func Test_winnerOfGame(t *testing.T) {
	type args struct {
		colors string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				colors: "AAABABB",
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				colors: "AA",
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				colors: "ABBBBBBBAAA",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerOfGame(tt.args.colors); got != tt.want {
				t.Errorf("winnerOfGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
