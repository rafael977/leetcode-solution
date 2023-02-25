package shufflethearray

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{2, 5, 1, 3, 4, 7},
				n:    3,
			},
			wantAns: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
				n:    4,
			},
			wantAns: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name: "tc1",
			args: args{
				nums: []int{1, 1, 2, 2},
				n:    2,
			},
			wantAns: []int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("shuffle() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
