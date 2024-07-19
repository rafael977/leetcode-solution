package leetcodesolution

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				matrix: [][]int{
					{3, 7, 8},
					{9, 11, 13},
					{15, 16, 17},
				},
			},
			wantAns: []int{15},
		},
		{
			name: "tc2",
			args: args{
				matrix: [][]int{
					{1, 10, 4, 2},
					{9, 3, 8, 7},
					{15, 16, 17, 12},
				},
			},
			wantAns: []int{12},
		},
		{
			name: "tc3",
			args: args{
				matrix: [][]int{
					{7, 8},
					{1, 2},
				},
			},
			wantAns: []int{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("luckyNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
