package countnodesequaltoaverageofsubtree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_averageOfSubtree(t *testing.T) {
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
				root: BuildTree("4,8,5,0,1,null,6"),
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1"),
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := averageOfSubtree(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("averageOfSubtree() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
