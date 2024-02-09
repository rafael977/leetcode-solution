package leetcodesolution

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
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
				nums: []int{1, 2, 3},
			},
			wantAns: []int{1, 2},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 2, 4, 8},
			},
			wantAns: []int{1, 2, 4, 8},
		},
		{
			name: "tc3",
			args: args{
				nums: []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720},
			},
			wantAns: []int{9, 18, 90, 180, 360, 720},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
