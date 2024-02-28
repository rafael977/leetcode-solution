package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_findBottomLeftValue(t *testing.T) {
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
				root: BuildTree("2,1,3"),
			},
			wantAns: 1,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,2,3,4,null,5,6,null,null,7"),
			},
			wantAns: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findBottomLeftValue(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("findBottomLeftValue() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
