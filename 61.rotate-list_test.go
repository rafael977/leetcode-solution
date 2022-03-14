package main

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				k:    2,
			},
			want: BuildLinkedList("4,5,1,2,3"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("0,1,2"),
				k:    4,
			},
			want: BuildLinkedList("2,0,1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
