package longeststringchain

import "testing"

func Test_longestStrChain(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				words: []string{"a", "b", "ba", "bca", "bda", "bdca"},
			},
			want: 4,
		},
		{
			name: "tc2",
			args: args{
				words: []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			},
			want: 5,
		},
		{
			name: "tc3",
			args: args{
				words: []string{"abcd", "dbqca"},
			},
			want: 1,
		},
		{
			name: "tc4",
			args: args{
				words: []string{"a", "ab", "ac", "bd", "abc", "abd", "abdd"},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain(tt.args.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
