package oddevenlinkedlist

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("1,2,3,4,5"),
			},
			want: BuildLinkedList("1,3,5,2,4"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("2,1,3,5,6,4,7"),
			},
			want: BuildLinkedList("2,3,6,7,1,5,4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); PrintLinkedList(got) != PrintLinkedList(tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
