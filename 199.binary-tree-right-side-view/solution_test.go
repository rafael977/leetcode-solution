package binarytreerightsideview

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("1,2,3,null,5,null,4"),
			},
			wantAns: []int{1, 3, 4},
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,null,3"),
			},
			wantAns: []int{1, 3},
		},
		{
			name: "tc3",
			args: args{
				root: BuildTree(""),
			},
			wantAns: []int{},
		},
		{
			name: "tc4",
			args: args{
				root: BuildTree("1,2"),
			},
			wantAns: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := rightSideView(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("rightSideView() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
