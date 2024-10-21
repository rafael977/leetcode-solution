package main

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_deleteDuplicates(t *testing.T) {
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
				head: BuildLinkedList("1,2,3,3,4,4,5"),
			},
			want: BuildLinkedList("1,2,5"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1,1,1,2,3"),
			},
			want: BuildLinkedList("2,3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
