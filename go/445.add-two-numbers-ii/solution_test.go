package addtwonumbersii

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "tc1",
			args: args{
				l1: BuildLinkedList("7,2,4,3"),
				l2: BuildLinkedList("5,6,4"),
			},
			want: BuildLinkedList("7,8,0,7"),
		},
		{
			name: "tc2",
			args: args{
				l1: BuildLinkedList("2,4,3"),
				l2: BuildLinkedList("5,6,4"),
			},
			want: BuildLinkedList("8,0,7"),
		},
		{
			name: "tc3",
			args: args{
				l1: BuildLinkedList("0"),
				l2: BuildLinkedList("0"),
			},
			want: BuildLinkedList("0"),
		},
		{
			name: "tc4",
			args: args{
				l1: BuildLinkedList("5"),
				l2: BuildLinkedList("5"),
			},
			want: BuildLinkedList("1,0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
