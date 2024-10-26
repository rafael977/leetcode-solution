// Created by Rafael Shen at 2024/10/26 09:37
// leetgo: 1.4.9
// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/

package heightofbinarytreeaftersubtreeremovalqueries

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_treeQueries(t *testing.T) {
	type args struct {
		root    *TreeNode
		queries []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				root:    BuildTree("1,3,4,2,null,6,5,null,null,null,null,null,7"),
				queries: []int{4},
			},
			wantAns: []int{2},
		},
		{
			name: "tc2",
			args: args{
				root:    BuildTree("5,8,9,2,1,3,7,4,6"),
				queries: []int{3, 2, 4, 8},
			},
			wantAns: []int{3, 2, 3, 2},
		},
		{
			name: "tc3",
			args: args{
				root:    BuildTree("1,null,5,3,null,2,4"),
				queries: []int{3, 5, 4, 2, 4},
			},
			wantAns: []int{1, 0, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := treeQueries(tt.args.root, tt.args.queries); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("treeQueries() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
