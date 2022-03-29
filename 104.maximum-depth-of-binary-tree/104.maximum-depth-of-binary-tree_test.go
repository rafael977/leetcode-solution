package main

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_maxDepth(t *testing.T) {
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
			want: 3,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,null,2"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
