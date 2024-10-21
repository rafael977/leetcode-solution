package partitionlist

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("1,4,3,2,5,2"),
				x:    3,
			},
			want: BuildLinkedList("1,2,2,4,3,5"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("2,1"),
				x:    2,
			},
			want: BuildLinkedList("1,2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
