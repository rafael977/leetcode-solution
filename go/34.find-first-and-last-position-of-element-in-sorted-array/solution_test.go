package findfirstandlastpositionofelementinsortedarray

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			wantAns: []int{3, 4},
		},
		{
			name: "tc2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			wantAns: []int{-1, -1},
		},
		{
			name: "tc3",
			args: args{
				nums:   []int{},
				target: 0,
			},
			wantAns: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("searchRange() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
