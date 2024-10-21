package setmismatch

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
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
				nums: []int{1, 2, 2, 4},
			},
			wantAns: []int{2, 3},
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 1},
			},
			wantAns: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findErrorNums(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findErrorNums() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
