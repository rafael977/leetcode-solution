package leetcodesolution

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "tc1",
			args: args{
				list1: BuildLinkedList("10,1,13,6,9,5"),
				list2: BuildLinkedList("1000000,1000001,1000002"),
				a:     3,
				b:     4,
			},
			want: BuildLinkedList("10,1,13,1000000,1000001,1000002,5"),
		},
		{
			name: "tc2",
			args: args{
				list1: BuildLinkedList("0,1,2,3,4,5,6"),
				list2: BuildLinkedList("1000000,1000001,1000002,1000003,1000004"),
				a:     2,
				b:     5,
			},
			want: BuildLinkedList("0,1,1000000,1000001,1000002,1000003,1000004,6"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
