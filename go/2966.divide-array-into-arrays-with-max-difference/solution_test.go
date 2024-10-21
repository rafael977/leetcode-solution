package dividearrayintoarrayswithmaxdifference

import (
	"reflect"
	"testing"
)

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
				k:    2,
			},
			wantAns: [][]int{
				{1, 1, 3}, {3, 4, 5}, {7, 8, 9},
			},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 3, 3, 2, 7, 3},
				k:    3,
			},
			wantAns: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := divideArray(tt.args.nums, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("divideArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
