package splitlinkedlistinparts

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []*ListNode
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("1,2,3"),
				k:    5,
			},
			wantAns: []*ListNode{
				BuildLinkedList("1"),
				BuildLinkedList("2"),
				BuildLinkedList("3"),
				BuildLinkedList(""),
				BuildLinkedList(""),
			},
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("1,2,3,4,5,6,7,8,9,10"),
				k:    3,
			},
			wantAns: []*ListNode{
				BuildLinkedList("1,2,3,4"),
				BuildLinkedList("5,6,7"),
				BuildLinkedList("8,9,10"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("splitListToParts() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
