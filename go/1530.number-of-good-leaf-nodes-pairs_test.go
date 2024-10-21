package leetcodesolution

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		root     *TreeNode
		distance int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				root:     BuildTree("1,2,3,null,4"),
				distance: 3,
			},
			wantAns: 1,
		},
		{
			name: "tc2",
			args: args{
				root:     BuildTree("1,2,3,4,5,6,7"),
				distance: 3,
			},
			wantAns: 2,
		},
		{
			name: "tc3",
			args: args{
				root:     BuildTree("7,1,4,6,null,5,3,null,null,null,null,null,2"),
				distance: 3,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countPairs(tt.args.root, tt.args.distance); gotAns != tt.wantAns {
				t.Errorf("countPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
