package insertinterval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			wantAns: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "tc2",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			wantAns: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("insert() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
