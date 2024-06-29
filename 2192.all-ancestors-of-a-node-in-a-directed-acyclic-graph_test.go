package leetcodesolution

import (
	"reflect"
	"testing"
)

func Test_getAncestors(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				n:     8,
				edges: [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}},
			},
			wantAns: [][]int{{}, {}, {}, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}},
		},
		{
			name: "tc2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			},
			wantAns: [][]int{{}, {0}, {0, 1}, {0, 1, 2}, {0, 1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := getAncestors(tt.args.n, tt.args.edges)
			for i := 0; i < tt.args.n; i++ {
				if !reflect.DeepEqual(gotAns[i], tt.wantAns[i]) {
					t.Errorf("getAncestors() = %v, want %v, ith = %v", gotAns, tt.wantAns, i)
					return
				}
			}

		})
	}
}
