package averageoflevelsinbinarytree

import (
	"reflect"
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []float64
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("3,9,20,null,null,15,7"),
			},
			wantAns: []float64{3, 14.5, 11},
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("3,9,20,15,7"),
			},
			wantAns: []float64{3, 14.5, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := averageOfLevels(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("averageOfLevels() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
