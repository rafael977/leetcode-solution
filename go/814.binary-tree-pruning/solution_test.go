package binarytreepruning

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_pruneTree(t *testing.T) {
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
			args: args{root: BuildTree("1,null,0,0,1")},
			want: BuildTree("1,null,0,null,1"),
		},
		{
			name: "tc2",
			args: args{root: BuildTree("1,0,1,0,0,0,1")},
			want: BuildTree("1,null,1,null,1"),
		},
		{
			name: "tc3",
			args: args{root: BuildTree("1,1,0,1,1,0,1,0")},
			want: BuildTree("1,1,0,1,1,null,1"),
		},
		{
			name: "tc4",
			args: args{root: BuildTree("0,null,0,0,0")},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
