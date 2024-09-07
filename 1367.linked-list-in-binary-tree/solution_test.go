package linkedlistinbinarytree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_isSubPath(t *testing.T) {
	type args struct {
		head *ListNode
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("4,2,8"),
				root: BuildTree("1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1,4,2,6"),
				root: BuildTree("1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3"),
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList("1,4,2,6,8"),
				root: BuildTree("1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3"),
			},
			want: false,
		},
		{
			name: "tc4",
			args: args{
				head: BuildLinkedList("1,10"),
				root: BuildTree("1,null,1,10,1,9"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubPath(tt.args.head, tt.args.root); got != tt.want {
				t.Errorf("isSubPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
