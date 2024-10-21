package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_removeNodes(t *testing.T) {
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
				head: BuildLinkedList("5,2,13,3,8"),
			},
			want: BuildLinkedList("13,8"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1,1,1,1"),
			},
			want: BuildLinkedList("1,1,1,1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNodes(tt.args.head); PrintLinkedList(tt.want) != PrintLinkedList(got) {
				t.Errorf("removeNodes() = %v, want %v",
					PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
