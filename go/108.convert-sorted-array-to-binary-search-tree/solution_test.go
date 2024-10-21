package convertsortedarraytobinarysearchtree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: BuildTree("0,-10,5,null,-3,null,9"),
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 3},
			},
			want: BuildTree("1,null,3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", PrintTree(got), PrintTree(tt.want))
			}
		})
	}
}
