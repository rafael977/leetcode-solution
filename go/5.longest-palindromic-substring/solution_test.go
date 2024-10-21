package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "tc1",
			args: args{
				s: "babad",
			},
			wantAns: "bab",
		},
		{
			name: "tc2",
			args: args{
				s: "cbbd",
			},
			wantAns: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestPalindrome(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("longestPalindrome() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
