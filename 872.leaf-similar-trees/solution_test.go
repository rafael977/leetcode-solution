package leafsimilartrees

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_leafSimilar(t *testing.T) {
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
				root1: BuildTree("3,5,1,6,2,9,8,null,null,7,4"),
				root2: BuildTree("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root1: BuildTree("1,2,3"),
				root2: BuildTree("1,3,2"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
