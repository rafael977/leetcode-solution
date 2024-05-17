package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_removeLeafNodes(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root:   BuildTree("1,2,3,2,null,2,4"),
				target: 2,
			},
			want: BuildTree("1,null,3,null,4"),
		},
		{
			name: "tc2",
			args: args{
				root:   BuildTree("1,3,3,3,2"),
				target: 3,
			},
			want: BuildTree("1,3,null,null,2"),
		},
		{
			name: "tc3",
			args: args{
				root:   BuildTree("1,2,null,2,null,2"),
				target: 2,
			},
			want: BuildTree("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLeafNodes(tt.args.root, tt.args.target); PrintTree(got) != PrintTree(tt.want) {
				t.Errorf("removeLeafNodes() = %v, want %v", PrintTree(got), PrintTree(tt.want))
			}
		})
	}
}
