package leetcodesolution

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			wantAns: []int{1},
		},
		{
			name: "tc2",
			args: args{
				n:     6,
				edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			},
			wantAns: []int{3, 4},
		},
		{
			name: "tc3",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			wantAns: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := findMinHeightTrees(tt.args.n, tt.args.edges)
			sort.Ints(gotAns)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findMinHeightTrees() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
