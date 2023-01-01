package allpathsfromsourcetotarget

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				graph: [][]int{{1, 2}, {3}, {3}, {}},
			},
			wantAns: [][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			name: "tc2",
			args: args{
				graph: [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			},
			wantAns: [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
