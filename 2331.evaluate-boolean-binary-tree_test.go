package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_evaluateTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("2,1,3,null,null,0,1"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("0"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateTree(tt.args.root); got != tt.want {
				t.Errorf("evaluateTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
