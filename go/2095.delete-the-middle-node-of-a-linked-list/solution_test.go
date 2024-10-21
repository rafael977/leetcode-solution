package deletethemiddlenodeofalinkedlist

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_deleteMiddle(t *testing.T) {
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
				head: BuildLinkedList("1,3,4,7,1,2,6"),
			},
			want: BuildLinkedList("1,3,4,1,2,6"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1,2,3,4"),
			},
			want: BuildLinkedList("1,2,4"),
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList("1,2"),
			},
			want: BuildLinkedList("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
