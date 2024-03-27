package leetcodesolution

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
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
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			wantAns: []int{2, 3},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 1, 2},
			},
			wantAns: []int{1},
		},
		{
			name: "tc3",
			args: args{
				nums: []int{},
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := findDuplicates(tt.args.nums)
			sort.Ints(gotAns)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findDuplicates() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
