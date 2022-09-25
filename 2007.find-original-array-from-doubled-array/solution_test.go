package findoriginalarrayfromdoubledarray

import (
	"reflect"
	"testing"
)

func Test_findOriginalArray(t *testing.T) {
	type args struct {
		changed []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				changed: []int{1, 3, 4, 2, 6, 8},
			},
			wantAns: []int{1, 3, 4},
		},
		{
			name: "tc2",
			args: args{
				changed: []int{6, 3, 0, 1},
			},
			wantAns: []int{},
		},
		{
			name: "tc3",
			args: args{
				changed: []int{1},
			},
			wantAns: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findOriginalArray(tt.args.changed); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findOriginalArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
