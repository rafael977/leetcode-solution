package breakapalindrome

import "testing"

func Test_breakPalindrome(t *testing.T) {
	type args struct {
		palindrome string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				palindrome: "abccba",
			},
			want: "aaccba",
		},
		{
			name: "tc2",
			args: args{
				palindrome: "a",
			},
			want: "",
		},
		{
			name: "tc3",
			args: args{
				palindrome: "aba",
			},
			want: "abb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakPalindrome(tt.args.palindrome); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
