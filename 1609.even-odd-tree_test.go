package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_isEvenOddTree(t *testing.T) {
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
				root: BuildTree("1,10,4,3,null,7,9,12,8,6,null,null,2"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("5,4,2,3,3,7"),
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("5,9,1,3,5,7"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.args.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
