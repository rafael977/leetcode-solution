package pacificatlanticwaterflow

import (
	"reflect"
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				heights: [][]int{
					{1, 2, 2, 3, 5},
					{3, 2, 3, 4, 4},
					{2, 4, 5, 3, 1},
					{6, 7, 1, 4, 5},
					{5, 1, 1, 2, 4},
				},
			},
			wantAns: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name: "tc2",
			args: args{
				heights: [][]int{
					{1},
				},
			},
			wantAns: [][]int{{0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := pacificAtlantic(tt.args.heights); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("pacificAtlantic() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
