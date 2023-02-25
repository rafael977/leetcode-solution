package invertbinarytree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_invertTree(t *testing.T) {
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
				root: BuildTree("4,2,7,1,3,6,9"),
			},
			want: BuildTree("4,7,2,9,6,3,1"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("2,1,3"),
			},
			want: BuildTree("2,3,1"),
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree(""),
			},
			want: BuildTree(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
