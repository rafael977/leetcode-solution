package findmodeinbinarysearchtree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_findMode(t *testing.T) {
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
				root: BuildTree("1,null,2,2"),
			},
			wantAns: []int{2},
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("0"),
			},
			wantAns: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMode(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findMode() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
