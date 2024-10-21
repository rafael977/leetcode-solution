package checkifastringcontainsallbinarycodesofsizek

import "testing"

func Test_hasAllCodes(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				s: "00110110",
				k: 2,
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				s: "0110",
				k: 1,
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				s: "0110",
				k: 2,
			},
			want: false,
		},
		{
			name: "tc4",
			args: args{
				s: "01100",
				k: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAllCodes(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("hasAllCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
