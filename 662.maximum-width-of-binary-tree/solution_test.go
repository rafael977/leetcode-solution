package maximumwidthofbinarytree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_widthOfBinaryTree(t *testing.T) {
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
				root: BuildTree("1,3,2,5,3,null,9"),
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,3,2,5,null,null,9,6,null,7"),
			},
			wantAns: 7,
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("1,3,2,5"),
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := widthOfBinaryTree(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("widthOfBinaryTree() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
