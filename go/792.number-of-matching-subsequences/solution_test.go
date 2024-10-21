package numberofmatchingsubsequences

import "testing"

func Test_numMatchingSubseq(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				s:     "abcde",
				words: []string{"a", "bb", "acd", "ace"},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				s:     "dsahjpjauf",
				words: []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numMatchingSubseq(tt.args.s, tt.args.words); gotAns != tt.wantAns {
				t.Errorf("numMatchingSubseq() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
