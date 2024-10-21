package checkcompletenessofabinarytree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_isCompleteTree(t *testing.T) {
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
				root: BuildTree("1,2,3,4,5,6"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,2,3,4,5,null,7"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCompleteTree(tt.args.root); got != tt.want {
				t.Errorf("isCompleteTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
