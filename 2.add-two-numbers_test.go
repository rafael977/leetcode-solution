package main

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
				l1: BuildLinkedList("2,4,3"),
				l2: BuildLinkedList("5,6,4"),
			},
			want: BuildLinkedList("7,0,8"),
		},
		{
			name: "tc2",
			args: args{
				l1: BuildLinkedList("0"),
				l2: BuildLinkedList("0"),
			},
			want: BuildLinkedList("0"),
		},
		{
			name: "tc3",
			args: args{
				l1: BuildLinkedList("2,4,6,1,1"),
				l2: BuildLinkedList("5,6,4"),
			},
			want: BuildLinkedList("7,0,1,2,1"),
		},
		{
			name: "tc4",
			args: args{
				l1: BuildLinkedList("9,9,9"),
				l2: BuildLinkedList("9,9"),
			},
			want: BuildLinkedList("8,9,0,1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
