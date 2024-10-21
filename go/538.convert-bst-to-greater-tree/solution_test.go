package convertbsttogreatertree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_convertBST(t *testing.T) {
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
				root: BuildTree("4,1,6,0,2,5,7,null,null,null,3,null,null,null,8"),
			},
			want: BuildTree("30,36,21,36,35,26,15,null,null,null,33,null,null,null,8"),
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("0,null,1"),
			},
			want: BuildTree("1,null,1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", PrintTree(got), PrintTree(tt.want))
			}
		})
	}
}
