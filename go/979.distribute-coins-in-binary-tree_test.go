package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_distributeCoins(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("3,0,0"),
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("0,3,0"),
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := distributeCoins(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("distributeCoins() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
