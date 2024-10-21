package convert1darrayinto2darray

import (
	"reflect"
	"testing"
)

func Test_construct2DArray(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				original: []int{1, 2, 3, 4},
				m:        2,
				n:        2,
			},
			wantAns: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "tc2",
			args: args{
				original: []int{1, 2, 3},
				m:        1,
				n:        3,
			},
			wantAns: [][]int{{1, 2, 3}},
		},
		{
			name: "tc3",
			args: args{
				original: []int{1, 2},
				m:        1,
				n:        1,
			},
			wantAns: [][]int(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("construct2DArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
