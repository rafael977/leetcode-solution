package main

import "testing"

func Test_scoreOfParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{
			name: "tc1",
			args: args{
				s: "()",
			},
			wantScore: 1,
		},
		{
			name: "tc2",
			args: args{
				s: "(())",
			},
			wantScore: 2,
		},
		{
			name: "tc3",
			args: args{
				s: "()()",
			},
			wantScore: 2,
		},
		{
			name: "tc4",
			args: args{
				s: "((())())",
			},
			wantScore: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := scoreOfParentheses(tt.args.s); gotScore != tt.wantScore {
				t.Errorf("scoreOfParentheses() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
