package topkfrequentelements

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			wantAns: []int{1, 2},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			wantAns: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("topKFrequent() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
