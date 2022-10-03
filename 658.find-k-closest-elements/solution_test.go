package findkclosestelements

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			wantAns: []int{1, 2, 3, 4},
		},
		{
			name: "tc2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			wantAns: []int{1, 2, 3, 4},
		},
		{
			name: "tc3",
			args: args{
				arr: []int{1, 3},
				k:   1,
				x:   2,
			},
			wantAns: []int{1},
		},
		{
			name: "tc4",
			args: args{
				arr: []int{1, 5, 10},
				k:   1,
				x:   4,
			},
			wantAns: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findClosestElements() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
