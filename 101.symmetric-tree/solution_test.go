package symmetrictree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_isSymmetric(t *testing.T) {
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
				root: BuildTree("1,2,2,3,4,4,3"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,2,2,null,3,null,3"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
