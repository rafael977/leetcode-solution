package linkedlistcycleii

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_detectCycle(t *testing.T) {
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
				head: GenCycle(BuildLinkedList("3,2,0,-4"), 1),
			},
			want: BuildLinkedList("2"),
		},
		{
			name: "tc2",
			args: args{
				head: GenCycle(BuildLinkedList("1,2"), 0),
			},
			want: BuildLinkedList("1"),
		},
		{
			name: "tc3",
			args: args{
				head: GenCycle(BuildLinkedList("1,2"), -1),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !(got == nil && tt.want == nil) && got.Val != tt.want.Val {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
