package verticalordertraversalofabinarytree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_verticalTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("3,9,20,null,null,15,7"),
			},
			wantAns: [][]int{
				{9}, {3, 15}, {20}, {7},
			},
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,2,3,4,5,6,7"),
			},
			wantAns: [][]int{
				{4}, {2}, {1, 5, 6}, {3}, {7},
			},
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("1,2,3,4,6,5,7"),
			},
			wantAns: [][]int{
				{4}, {2}, {1, 5, 6}, {3}, {7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := verticalTraversal(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("verticalTraversal() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
