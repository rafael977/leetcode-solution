package findandreplacepattern

import (
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	type args struct {
		words   []string
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "tc1",
			args: args{
				words:   []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
				pattern: "abb",
			},
			wantAns: []string{"mee", "aqq"},
		},
		{
			name: "tc2",
			args: args{
				words:   []string{"a", "b", "c"},
				pattern: "a",
			},
			wantAns: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAndReplacePattern(tt.args.words, tt.args.pattern); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAndReplacePattern() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
