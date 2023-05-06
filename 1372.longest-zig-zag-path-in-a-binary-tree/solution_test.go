package longestzigzagpathinabinarytree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_longestZigZag(t *testing.T) {
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
				root: BuildTree("1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1"),
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,1,1,null,1,null,null,1,1,null,1"),
			},
			wantAns: 4,
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("1"),
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestZigZag(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("longestZigZag() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
