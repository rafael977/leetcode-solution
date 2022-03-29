package main

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_sortList(t *testing.T) {
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
				head: BuildLinkedList("4,2,1,3"),
			},
			want: BuildLinkedList("1,2,3,4"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("-1,5,3,4,0"),
			},
			want: BuildLinkedList("-1,0,3,4,5"),
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList(""),
			},
			want: BuildLinkedList(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
