package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_balanceBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("1,null,2,null,3,null,4,null,null"),
			},
			want: BuildTree("2,1,3,null,null,null,4"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("2,1,3"),
			},
			want: BuildTree("2,1,3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceBST(tt.args.root); PrintTree(got) != PrintTree(tt.want) {
				t.Errorf("balanceBST() = %v, want %v", PrintTree(got), PrintTree(tt.want))
			}
		})
	}
}
