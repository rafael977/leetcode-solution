package twosumivinputisabst

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("5,3,6,2,4,null,7"),
				k:    9,
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("5,3,6,2,4,null,7"),
				k:    28,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
