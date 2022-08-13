package longestvalidparentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "tc1",
			args:    args{s: "(()"},
			wantAns: 2,
		},
		{
			name:    "tc2",
			args:    args{s: ")()())"},
			wantAns: 4,
		},
		{
			name:    "tc3",
			args:    args{s: ""},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestValidParentheses(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("longestValidParentheses() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
