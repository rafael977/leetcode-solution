package trimabinarysearchtree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_trimBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("1,0,2"),
				low:  1,
				high: 2,
			},
			want: BuildTree("1,null,2"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("3,0,4,null,2,null,null,1"),
				low:  1,
				high: 3,
			},
			want: BuildTree("3,2,null,1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
