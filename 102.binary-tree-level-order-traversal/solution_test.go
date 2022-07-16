package binarytreelevelordertraversal

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("3,9,20,null,null,15,7"),
			},
			wantAns: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("3"),
			},
			wantAns: [][]int{{3}},
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree(""),
			},
			wantAns: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := levelOrder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("levelOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
