package minimumnumberofverticestoreachallnodes

import (
	"reflect"
	"testing"
)

func Test_findSmallestSetOfVertices(t *testing.T) {
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
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}},
			},
			wantAns: []int{0, 3},
		},
		{
			name: "tc2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}},
			},
			wantAns: []int{0, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findSmallestSetOfVertices(tt.args.n, tt.args.edges); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findSmallestSetOfVertices() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
