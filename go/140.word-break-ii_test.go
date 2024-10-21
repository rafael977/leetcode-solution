package leetcodesolution

import (
	"reflect"
	"sort"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "tc1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			wantAns: []string{"cat sand dog", "cats and dog"},
		},
		{
			name: "tc2",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			wantAns: []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		},
		{
			name: "tc3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			wantAns: []string(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := wordBreak(tt.args.s, tt.args.wordDict)
			sort.Strings(gotAns)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("wordBreak() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
