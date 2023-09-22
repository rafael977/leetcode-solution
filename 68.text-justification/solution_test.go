package textjustification

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "tc1",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name: "tc2",
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name: "tc3",
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20,
			},
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
		{
			name: "tc4",
			args: args{
				words:    []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"},
				maxWidth: 16,
			},
			want: []string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
