package findthedifferenceoftwoarrays

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			wantAns: [][]int{{1, 3}, {4, 6}},
		},
		{
			name: "tc2",
			args: args{
				nums1: []int{1, 2, 3, 3},
				nums2: []int{1, 1, 2, 2},
			},
			wantAns: [][]int{{3}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := findDifference(tt.args.nums1, tt.args.nums2)
			for i := range gotAns {
				sort.Ints(gotAns[i])
			}
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
