package sumofevennumbersafterqueries

import (
	"reflect"
	"testing"
)

func Test_sumEvenAfterQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				nums:    []int{1, 2, 3, 4},
				queries: [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}},
			},
			wantAns: []int{8, 6, 2, 4},
		},
		{
			name: "tc2",
			args: args{
				nums:    []int{1},
				queries: [][]int{{4, 0}},
			},
			wantAns: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := sumEvenAfterQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("sumEvenAfterQueries() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
