package recoverbinarysearchtree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("1,3,null,null,2"),
			},
			want: BuildTree("3,1,null,null,2"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("3,1,4,null,null,2"),
			},
			want: BuildTree("2,1,4,null,null,3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.args.root
			recoverTree(root)
			if !reflect.DeepEqual(root, tt.want) {
				t.Errorf("recoverTree(): %v, want: %v", PrintTree(root), PrintTree(tt.want))
			}
		})
	}
}
