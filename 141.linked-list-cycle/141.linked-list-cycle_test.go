package main

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_hasCycle(t *testing.T) {
	linkNode := func(head *ListNode, pos int) *ListNode {
		tail := head
		for tail.Next != nil {
			tail = tail.Next
		}
		cur := head
		for i := 0; i < pos; i++ {
			cur = cur.Next
		}
		tail.Next = cur

		return head
	}

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
				head: linkNode(BuildLinkedList("3,2,0,-4"), 1),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				head: linkNode(BuildLinkedList("1,2"), 0),
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList("1"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
