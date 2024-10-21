package pseudopalindromicpathsinabinarytree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_pseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "tc1",
			args:    args{root: BuildTree("2,3,1,3,1,null,1")},
			wantAns: 2,
		},
		{
			name:    "tc2",
			args:    args{root: BuildTree("2,1,1,1,3,null,null,null,null,null,1")},
			wantAns: 1,
		},
		{
			name:    "tc3",
			args:    args{root: BuildTree("9")},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := pseudoPalindromicPaths(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
