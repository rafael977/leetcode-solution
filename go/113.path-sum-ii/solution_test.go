package pathsumii

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				root:      BuildTree("5,4,8,11,null,13,4,7,2,null,null,5,1"),
				targetSum: 22,
			},
			wantAns: [][]int{
				{5, 4, 11, 2}, {5, 8, 4, 5},
			},
		},
		{
			name: "tc2",
			args: args{
				root:      BuildTree("1,2,3"),
				targetSum: 5,
			},
			wantAns: nil,
		},
		{
			name: "tc3",
			args: args{
				root:      BuildTree("1"),
				targetSum: 0,
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("pathSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
