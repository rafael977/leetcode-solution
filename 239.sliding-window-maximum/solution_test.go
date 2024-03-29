package slidingwindowmaximum

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
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
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			wantAns: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, -1},
				k:    1,
			},
			wantAns: []int{1, -1},
		},
		{
			name: "tc3",
			args: args{
				nums: []int{7, 2, 4},
				k:    2,
			},
			wantAns: []int{7, 4},
		},
		{
			name: "tc4",
			args: args{
				nums: []int{9, 10, 9, -7, -4, -8, 2, -6},
				k:    5,
			},
			wantAns: []int{10, 10, 9, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("maxSlidingWindow() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
