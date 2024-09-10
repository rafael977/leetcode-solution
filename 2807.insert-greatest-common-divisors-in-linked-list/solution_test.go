package insertgreatestcommondivisorsinlinkedlist

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_insertGreatestCommonDivisors(t *testing.T) {
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
				head: BuildLinkedList("18,6,10,3"),
			},
			want: BuildLinkedList("18,6,6,2,10,1,3"),
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("7"),
			},
			want: BuildLinkedList("7"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertGreatestCommonDivisors(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
