package lowestcommonancestorofabinarysearchtree

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
		want int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("6,2,8,0,4,7,9,null,null,3,5"),
				p:    BuildTree("2"),
				q:    BuildTree("8"),
			},
			want: 6,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("6,2,8,0,4,7,9,null,null,3,5"),
				p:    BuildTree("2"),
				q:    BuildTree("4"),
			},
			want: 2,
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("2,1"),
				p:    BuildTree("2"),
				q:    BuildTree("1"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); got.Val != tt.want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
