package addtoarrayformofinteger

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		num []int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				num: []int{1, 2, 0, 0},
				k:   34,
			},
			wantAns: []int{1, 2, 3, 4},
		},
		{
			name: "tc2",
			args: args{
				num: []int{2, 7, 4},
				k:   181,
			},
			wantAns: []int{4, 5, 5},
		},
		{
			name: "tc3",
			args: args{
				num: []int{2, 1, 5},
				k:   806,
			},
			wantAns: []int{1, 0, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := addToArrayForm(tt.args.num, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("addToArrayForm() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
