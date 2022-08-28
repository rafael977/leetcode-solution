package palindromelinkedlist

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("1,2,2,1"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1,2"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
