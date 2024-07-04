package leetcodesolution

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_mergeNodes(t *testing.T) {
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
				head: BuildLinkedList("0,3,1,0,4,5,2,0"),
			},
			want: BuildLinkedList("4,11"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("0,1,0,3,0,2,2,0"),
			},
			want: BuildLinkedList("1,3,4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
