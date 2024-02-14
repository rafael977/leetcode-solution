package leetcodesolution

import (
	"reflect"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
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
				nums: []int{3, 1, -2, -5, 2, -4},
			},
			wantAns: []int{3, -2, 1, -5, 2, -4},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{-1, 1},
			},
			wantAns: []int{1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := rearrangeArray(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("rearrangeArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
