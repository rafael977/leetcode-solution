package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_sumNumbers(t *testing.T) {
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
				root: BuildTree("1,2,3"),
			},
			wantAns: 25,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("4,9,0,5,1"),
			},
			wantAns: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := sumNumbers(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("sumNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
