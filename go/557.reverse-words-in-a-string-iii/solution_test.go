package reversewordsinastringiii

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{s: "Let's take LeetCode contest"},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "tc2",
			args: args{s: "God Ding"},
			want: "doG gniD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
