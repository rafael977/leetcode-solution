// Created by Rafael Shen at 2024/10/21 12:54
// leetgo: 1.4.9
// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/

package main

import "testing"

func Test_maxUniqueSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				s: "ababccc",
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				s: "aba",
			},
			wantAns: 2,
		},
		{
			name: "tc3",
			args: args{
				s: "aaaaa",
			},
			wantAns: 2,
		},
		{
			name: "tc4",
			args: args{
				s: "wwwzfvedwfvhsww",
			},
			wantAns: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxUniqueSplit(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxUniqueSplit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
