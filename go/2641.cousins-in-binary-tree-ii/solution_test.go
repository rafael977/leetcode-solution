// Created by Rafael Shen at 2024/10/23 09:52
// leetgo: 1.4.9
// https://leetcode.com/problems/cousins-in-binary-tree-ii/

package cousinsinbinarytreeii

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_replaceValueInTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("5,4,9,1,10,null,7"),
			},
			wantAns: BuildTree("0,0,0,7,7,null,11"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("3,1,2"),
			},
			wantAns: BuildTree("0,0,0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := replaceValueInTree(tt.args.root); PrintTree(gotAns) != PrintTree(tt.wantAns) {
				t.Errorf("replaceValueInTree() = %v, want %v", PrintTree(gotAns), PrintTree(tt.wantAns))
			}
		})
	}
}
