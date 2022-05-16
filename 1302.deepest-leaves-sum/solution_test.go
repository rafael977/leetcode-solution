package deepestleavessum

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_deepestLeavesSum(t *testing.T) {
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
				root: BuildTree("1,2,3,4,5,null,6,7,null,null,null,null,8"),
			},
			want: 15,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("6,7,8,2,7,1,3,9,null,1,4,null,null,null,5"),
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepestLeavesSum(tt.args.root); got != tt.want {
				t.Errorf("deepestLeavesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
