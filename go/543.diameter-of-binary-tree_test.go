package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("1,2,3,4,5"),
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,2"),
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := diameterOfBinaryTree(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
