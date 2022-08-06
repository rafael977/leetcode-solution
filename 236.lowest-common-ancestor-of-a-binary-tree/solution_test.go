package lowestcommonancestorofabinarytree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("3,5,1,6,2,0,8,null,null,7,4"),
				p:    BuildTree("5"),
				q:    BuildTree("1"),
			},
			want: BuildTree("3"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("3,5,1,6,2,0,8,null,null,7,4"),
				p:    BuildTree("5"),
				q:    BuildTree("4"),
			},
			want: BuildTree("5"),
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("1,2"),
				p:    BuildTree("1"),
				q:    BuildTree("2"),
			},
			want: BuildTree("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
