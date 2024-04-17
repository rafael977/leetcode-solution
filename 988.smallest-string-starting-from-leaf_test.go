package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_smallestFromLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("0,1,2,3,4,3,4"),
			},
			wantAns: "dba",
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("25,1,3,1,3,0,2"),
			},
			wantAns: "adz",
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("2,2,1,null,1,0,null,0"),
			},
			wantAns: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := smallestFromLeaf(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("smallestFromLeaf() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
