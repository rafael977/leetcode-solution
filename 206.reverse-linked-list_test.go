package leetcodesolution

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_reverseList(t *testing.T) {
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
				head: BuildLinkedList("1,2,3,4,5"),
			},
			want: BuildLinkedList("5,4,3,2,1"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1,2"),
			},
			want: BuildLinkedList("2,1"),
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
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
