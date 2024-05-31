package leetcodesolution

import (
	"reflect"
	"sort"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 2, 1, 3, 2, 5},
			},
			wantAns: []int{3, 5},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{-1, 0},
			},
			wantAns: []int{-1, 0},
		},
		{
			name: "tc3",
			args: args{
				nums: []int{0, 1},
			},
			wantAns: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := singleNumber(tt.args.nums)
			sort.Ints(gotAns)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("singleNumber() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
