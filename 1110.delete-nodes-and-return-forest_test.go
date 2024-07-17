package leetcodesolution

import (
	"sort"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_delNodes(t *testing.T) {
	type args struct {
		root      *TreeNode
		to_delete []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []*TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root:      BuildTree("1,2,3,4,5,6,7"),
				to_delete: []int{3, 5},
			},
			wantAns: []*TreeNode{
				BuildTree("1,2,null,4"),
				BuildTree("6"),
				BuildTree("7"),
			},
		},
		{
			name: "tc2",
			args: args{
				root:      BuildTree("1,2,4,null,3"),
				to_delete: []int{3, 5},
			},
			wantAns: []*TreeNode{
				BuildTree("1,2,4"),
			},
		},
		{
			name: "tc3",
			args: args{
				root:      BuildTree("1,2,null,4,3"),
				to_delete: []int{2, 3},
			},
			wantAns: []*TreeNode{
				BuildTree("1"),
				BuildTree("4"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := delNodes(tt.args.root, tt.args.to_delete)
			if len(gotAns) != len(tt.wantAns) {
				t.Errorf("length wrong: delNodes() = %v, want %v",
					len(gotAns), len(tt.wantAns))
				return
			}
			sort.Slice(gotAns, func(i, j int) bool {
				return gotAns[i].Val < gotAns[j].Val
			})
			for i := range gotAns {
				if PrintTree(gotAns[i]) != PrintTree(tt.wantAns[i]) {
					t.Errorf("delNodes() = %v, want %v",
						PrintTree(gotAns[i]), PrintTree(tt.wantAns[i]))
					return
				}
			}
		})
	}
}
