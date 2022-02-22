package main

import (
	"testing"
)

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); gotAns != tt.wantAns {
				t.Errorf("ladderLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
