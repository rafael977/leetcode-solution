package numberofnodesinthesubtreewiththesamelabel

import (
	"reflect"
	"testing"
)

func Test_countSubTrees(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		labels string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				n:      7,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
				labels: "abaedcd",
			},
			wantAns: []int{2, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "tc2",
			args: args{
				n:      4,
				edges:  [][]int{{0, 1}, {1, 2}, {0, 3}},
				labels: "bbbb",
			},
			wantAns: []int{4, 2, 1, 1},
		},
		{
			name: "tc3",
			args: args{
				n:      5,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}},
				labels: "aabab",
			},
			wantAns: []int{3, 2, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSubTrees(tt.args.n, tt.args.edges, tt.args.labels); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("countSubTrees() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
