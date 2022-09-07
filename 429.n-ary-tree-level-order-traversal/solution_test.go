package narytreelevelordertraversal

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				root: buildNTree("1,null,3,2,4,null,5,6"),
			},
			wantAns: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			name: "tc2",
			args: args{
				root: buildNTree("1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14"),
			},
			wantAns: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := levelOrder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("levelOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
