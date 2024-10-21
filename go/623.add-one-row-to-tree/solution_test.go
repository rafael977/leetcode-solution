package addonerowtotree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root:  BuildTree("4,2,6,3,1,5"),
				val:   1,
				depth: 2,
			},
			want: BuildTree("4,1,1,2,null,null,6,3,1,5"),
		},
		{
			name: "tc2",
			args: args{
				root:  BuildTree("4,2,null,3,1"),
				val:   1,
				depth: 3,
			},
			want: BuildTree("4,2,null,1,1,3,null,null,1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", PrintTree(got), PrintTree(tt.want))
			}
		})
	}
}
