package largestsubstringbetweentwoequalcharacters

import "testing"

func Test_maxLengthBetweenEqualCharacters(t *testing.T) {
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
				s: "aa",
			},
			want: 0,
		},
		{
			name: "tc2",
			args: args{
				s: "abca",
			},
			want: 2,
		},
		{
			name: "tc3",
			args: args{
				s: "cbzxy",
			},
			want: -1,
		},
		{
			name: "tc4",
			args: args{
				s: "mgntdygtxrvxjnwksqhxuxtrv",
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLengthBetweenEqualCharacters(tt.args.s); got != tt.want {
				t.Errorf("maxLengthBetweenEqualCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
