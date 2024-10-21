package maximumnumberofvowelsinasubstringofgivenlength

import "testing"

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				s: "abciiidef",
				k: 3,
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				s: "aeiou",
				k: 2,
			},
			wantAns: 2,
		},
		{
			name: "tc3",
			args: args{
				s: "leetcode",
				k: 3,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxVowels(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxVowels() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
