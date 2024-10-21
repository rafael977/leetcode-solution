package leetcodesolution

import "testing"

func Test_reverseParentheses(t *testing.T) {
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
			args: args{
				s: "(abcd)",
			},
			want: "dcba",
		},
		{
			name: "tc2",
			args: args{
				s: "(u(love)i)",
			},
			want: "iloveu",
		},
		{
			name: "tc3",
			args: args{
				s: "(ed(et(oc))el)",
			},
			want: "leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
