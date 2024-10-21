package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_sumOfLeftLeaves(t *testing.T) {
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
				root: BuildTree("3,9,20,null,null,15,7"),
			},
			wantAns: 24,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1"),
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := sumOfLeftLeaves(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
