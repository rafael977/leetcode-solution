package spiralmatrixiv

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_spiralMatrix(t *testing.T) {
	type args struct {
		m    int
		n    int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "tc1",
			args: args{
				m:    3,
				n:    5,
				head: BuildLinkedList("3,0,2,6,8,1,7,9,4,2,5,5,0"),
			},
			want: [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}},
		},
		{
			name: "tc2",
			args: args{
				m:    1,
				n:    4,
				head: BuildLinkedList("0,1,2"),
			},
			want: [][]int{{0, 1, 2, -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrix(tt.args.m, tt.args.n, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
