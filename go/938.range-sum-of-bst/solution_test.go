package rangesumofbst

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("10,5,15,3,7,null,18"),
				low:  7,
				high: 15,
			},
			want: 32,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("10,5,15,3,7,13,18,1,null,6"),
				low:  6,
				high: 10,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
