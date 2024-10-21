package minimumdepthofbinarytree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("3,9,20,null,null,15,7"),
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("2,null,3,null,4,null,5,null,6"),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
