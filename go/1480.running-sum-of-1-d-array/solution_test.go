package runningsumof1darray

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name   string
		args   args
		wantRs []int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			wantRs: []int{1, 3, 6, 10},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			wantRs: []int{1, 2, 3, 4, 5},
		},
		{
			name: "tc3",
			args: args{
				nums: []int{3, 1, 2, 10, 1},
			},
			wantRs: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRs := runningSum(tt.args.nums); !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("runningSum() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}
