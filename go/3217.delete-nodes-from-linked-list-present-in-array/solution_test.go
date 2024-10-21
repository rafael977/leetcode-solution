package deletenodesfromlinkedlistpresentinarray

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_modifiedList(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 3},
				head: BuildLinkedList("1,2,3,4,5"),
			},
			want: BuildLinkedList("4,5"),
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1},
				head: BuildLinkedList("1,2,1,2,1,2"),
			},
			want: BuildLinkedList("2,2,2"),
		},
		{
			name: "tc3",
			args: args{
				nums: []int{5},
				head: BuildLinkedList("1,2,3,4"),
			},
			want: BuildLinkedList("1,2,3,4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedList(tt.args.nums, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedList() = %v, want %v", PrintLinkedList(got), PrintLinkedList(tt.want))
			}
		})
	}
}
