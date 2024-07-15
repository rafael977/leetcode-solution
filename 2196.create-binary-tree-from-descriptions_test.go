package leetcodesolution

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_createBinaryTree(t *testing.T) {
	type args struct {
		descriptions [][]int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "tc1",
			args: args{
				descriptions: [][]int{
					{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1},
				},
			},
			want: BuildTree("50,20,80,15,17,19"),
		},
		{
			name: "tc2",
			args: args{
				descriptions: [][]int{
					{1, 2, 1}, {2, 3, 0}, {3, 4, 1},
				},
			},
			want: BuildTree("1,2,null,null,3,4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBinaryTree(tt.args.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
