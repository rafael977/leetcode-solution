package leetcodesolution

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_nodesBetweenCriticalPoints(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("3,1"),
			},
			want: []int{-1, -1},
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("5,3,1,2,5,1,2"),
			},
			want: []int{1, 3},
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList("1,3,2,2,3,2,2,2,7"),
			},
			want: []int{3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodesBetweenCriticalPoints(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodesBetweenCriticalPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
