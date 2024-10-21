package ranktransformofanarray

import (
	"reflect"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				arr: []int{40, 10, 20, 30},
			},
			wantAns: []int{4, 1, 2, 3},
		},
		{
			name: "tc2",
			args: args{
				arr: []int{100, 100, 100},
			},
			wantAns: []int{1, 1, 1},
		},
		{
			name: "tc3",
			args: args{
				arr: []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			},
			wantAns: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := arrayRankTransform(tt.args.arr); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("arrayRankTransform() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
