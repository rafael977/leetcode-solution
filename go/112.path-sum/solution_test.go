package pathsum

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				root:      BuildTree("5,4,8,11,null,13,4,7,2,null,null,null,1"),
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				root:      BuildTree("1,2,3"),
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				root:      BuildTree(""),
				targetSum: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
