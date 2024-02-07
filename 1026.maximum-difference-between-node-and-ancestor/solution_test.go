package maximumdifferencebetweennodeandancestor

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_maxAncestorDiff(t *testing.T) {
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
				root: BuildTree("8,3,10,1,6,null,14,null,null,4,7,13"),
			},
			wantAns: 7,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,null,2,null,0,3"),
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxAncestorDiff(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("maxAncestorDiff() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
