// Created by Rafael Shen at 2024/10/24 10:51
// leetgo: 1.4.9
// https://leetcode.com/problems/flip-equivalent-binary-trees/

package flipequivalentbinarytrees

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_flipEquiv(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				root1: BuildTree("1,2,3,4,5,6,null,null,null,7,8"),
				root2: BuildTree("1,3,2,null,6,4,5,null,null,null,null,8,7"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root1: BuildTree(""),
				root2: BuildTree(""),
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				root1: BuildTree(""),
				root2: BuildTree("1"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipEquiv(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
