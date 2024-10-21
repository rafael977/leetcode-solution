package flattenbinarytreetolinkedlist

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_flatten(t *testing.T) {
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
				root: BuildTree("1,2,5,3,4,null,6"),
			},
			want: BuildTree("1,null,2,null,3,null,4,null,5,null,6"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree(""),
			},
			want: BuildTree(""),
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree("0"),
			},
			want: BuildTree("0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("want: %v, got: %v", PrintTree(tt.want), PrintTree(tt.args.root))
			}
		})
	}
}
