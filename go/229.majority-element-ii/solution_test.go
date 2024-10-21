package majorityelementii

import (
	"reflect"
	"sort"
	"testing"
)

func Test_majorityElement(t *testing.T) {
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
				nums: []int{3, 2, 3},
			},
			wantAns: []int{3},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1},
			},
			wantAns: []int{1},
		},
		{
			name: "tc3",
			args: args{
				nums: []int{1, 2},
			},
			wantAns: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := majorityElement(tt.args.nums)
			sort.Ints(gotAns)
			if !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("majorityElement() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
