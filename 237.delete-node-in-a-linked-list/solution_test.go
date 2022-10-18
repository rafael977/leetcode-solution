package deletenodeinalinkedlist

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		node func(*ListNode) *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("4,5,1,9"),
				node: func(ln *ListNode) *ListNode { return ln.Next },
			},
			want: BuildLinkedList("4,1,9"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("4,5,1,9"),
				node: func(ln *ListNode) *ListNode { return ln.Next.Next },
			},
			want: BuildLinkedList("4,5,9"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node(tt.args.head))
			if PrintLinkedList(tt.want) != PrintLinkedList(tt.args.head) {
				t.Errorf("Expected: %v, got: %v", PrintLinkedList(tt.want),
					PrintLinkedList(tt.args.head))
			}
		})
	}
}
