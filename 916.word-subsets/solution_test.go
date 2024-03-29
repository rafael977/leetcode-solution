package wordsubsets

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "tc1",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "o"},
			},
			wantAns: []string{"facebook", "google", "leetcode"},
		},
		{
			name: "tc2",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"l", "e"},
			},
			wantAns: []string{"apple", "google", "leetcode"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := wordSubsets(tt.args.words1, tt.args.words2); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("wordSubsets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
