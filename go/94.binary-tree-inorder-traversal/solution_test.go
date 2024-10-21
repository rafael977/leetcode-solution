package binarytreeinordertraversal

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tc1",
			args: args{root: BuildTree("1,null,2,3")},
			want: []int{1, 3, 2},
		},
		{
			name: "tc2",
			args: args{root: BuildTree("1,2,3,4,5,6,7,8")},
			want: []int{8, 4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
