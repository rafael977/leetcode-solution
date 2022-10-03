package removenthnodefromendoflist

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
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
				n:    2,
			},
			want: BuildLinkedList("1,2,3,5"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1"),
				n:    1,
			},
			want: nil,
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList("1,2"),
				n:    1,
			},
			want: BuildLinkedList("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
