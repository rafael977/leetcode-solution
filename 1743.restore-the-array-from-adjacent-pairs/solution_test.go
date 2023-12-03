package restorethearrayfromadjacentpairsgo

import (
	"reflect"
	"slices"
	"testing"
)

func Test_restoreArray(t *testing.T) {
	type args struct {
		adjacentPairs [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				adjacentPairs: [][]int{{2, 1}, {3, 4}, {3, 2}},
			},
			wantAns: []int{1, 2, 3, 4},
		},
		{
			name: "tc2",
			args: args{
				adjacentPairs: [][]int{{4, -2}, {1, 4}, {-3, 1}},
			},
			wantAns: []int{-2, 4, 1, -3},
		},
		{
			name: "tc3",
			args: args{
				adjacentPairs: [][]int{{100000, -100000}},
			},
			wantAns: []int{100000, -100000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rev := make([]int, len(tt.wantAns))
			copy(rev, tt.wantAns)
			slices.Reverse(rev)
			if gotAns := restoreArray(tt.args.adjacentPairs); !reflect.DeepEqual(gotAns, tt.wantAns) && !reflect.DeepEqual(gotAns, rev) {
				t.Errorf("restoreArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
