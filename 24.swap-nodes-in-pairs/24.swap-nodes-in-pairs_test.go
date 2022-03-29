package main

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_swapPairs(t *testing.T) {
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
				head: BuildLinkedList("1,2,3,4"),
			},
			want: BuildLinkedList("2,1,4,3"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1"),
			},
			want: BuildLinkedList("1"),
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList(""),
			},
			want: BuildLinkedList(""),
		},
		{
			name: "tc4",
			args: args{
				head: BuildLinkedList("1,2,3,4,5"),
			},
			want: BuildLinkedList("2,1,4,3,5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Logf("got = %v", PrintLinkedList(got))
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
