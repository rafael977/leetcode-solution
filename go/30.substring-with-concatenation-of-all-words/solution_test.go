package substringwithconcatenationofallwords

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			wantAns: []int{0, 9},
		},
		{
			name: "tc2",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			wantAns: nil,
		},
		{
			name: "tc3",
			args: args{
				s:     "barfoofoobarthefoobarman",
				words: []string{"bar", "foo", "the"},
			},
			wantAns: []int{6, 9, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
