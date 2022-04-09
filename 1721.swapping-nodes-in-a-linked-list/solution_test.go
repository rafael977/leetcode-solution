package swappingnodesinalinkedlist

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_swapNodes(t *testing.T) {
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
			want: BuildLinkedList("1,4,3,2,5"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("7,9,6,6,7,8,3,0,9,5"),
				k:    5,
			},
			want: BuildLinkedList("7,9,6,6,8,7,3,0,9,5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
