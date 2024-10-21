package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_doubleIt(t *testing.T) {
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
				head: BuildLinkedList("1,8,9"),
			},
			want: BuildLinkedList("3,7,8"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("9,9,9"),
			},
			want: BuildLinkedList("1,9,9,8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleIt(tt.args.head); PrintLinkedList(got) != PrintLinkedList(tt.want) {
				t.Errorf("doubleIt() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
