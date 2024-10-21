package combinationsumiii

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name:    "tc1",
			args:    args{k: 3, n: 7},
			wantAns: [][]int{{1, 2, 4}},
		},
		{
			name:    "tc2",
			args:    args{k: 3, n: 9},
			wantAns: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name:    "tc3",
			args:    args{k: 4, n: 1},
			wantAns: [][]int{},
		},
		{
			name:    "tc4",
			args:    args{k: 8, n: 36},
			wantAns: [][]int{{1, 2, 3, 4, 5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum3() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
