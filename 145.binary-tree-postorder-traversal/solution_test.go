package binarytreepostordertraversal

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("1,null,2,3"),
			},
			wantAns: []int{3, 2, 1},
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree(""),
			},
			wantAns: []int(nil),
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("1"),
			},
			wantAns: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := postorderTraversal(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("postorderTraversal() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
