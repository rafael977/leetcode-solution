package validatebinarysearchtree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_isValidBST(t *testing.T) {
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
				root: BuildTree("2,1,3"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("5,1,4,null,null,3,6"),
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("2,2,2"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
