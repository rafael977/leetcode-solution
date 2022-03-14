package main

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
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
				list1: BuildLinkedList("1,2,4"),
				list2: BuildLinkedList("1,3,4"),
			},
			want: BuildLinkedList("1,1,2,3,4,4"),
		},
		{
			name: "tc2",
			args: args{
				list1: BuildLinkedList(""),
				list2: BuildLinkedList(""),
			},
			want: BuildLinkedList(""),
		},
		{
			name: "tc3",
			args: args{
				list1: BuildLinkedList(""),
				list2: BuildLinkedList("0"),
			},
			want: BuildLinkedList("0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
