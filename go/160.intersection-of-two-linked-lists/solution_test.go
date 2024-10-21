package intersectionoftwolinkedlists

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "tc1",
			args: args{
				headA: BuildLinkedList("4,1"),
				headB: BuildLinkedList("5,6,1"),
			},
			want: BuildLinkedList("8,4,5"),
		},
		{
			name: "tc2",
			args: args{
				headA: BuildLinkedList("1,9,1"),
				headB: BuildLinkedList("3"),
			},
			want: BuildLinkedList("2,4"),
		},
		{
			name: "tc3",
			args: args{
				headA: BuildLinkedList("2,6,4"),
				headB: BuildLinkedList("1,5"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(linkList(tt.args.headA, tt.want), linkList(tt.args.headB, tt.want)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func linkList(head, intersection *ListNode) *ListNode {
	t := head
	for t.Next != nil {
		t = t.Next
	}
	t.Next = intersection
	return head
}
