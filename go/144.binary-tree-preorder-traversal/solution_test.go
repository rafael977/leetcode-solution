package binarytreepreordertraversal

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_preorderTraversal(t *testing.T) {
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
			wantAns: []int{1, 2, 3},
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree(""),
			},
			wantAns: nil,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1"),
			},
			wantAns: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := preorderTraversal(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("preorderTraversal() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
