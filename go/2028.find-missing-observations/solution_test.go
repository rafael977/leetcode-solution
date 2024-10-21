package findmissingobservations

import (
	"reflect"
	"testing"
)

func Test_missingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				rolls: []int{3, 2, 4, 3},
				mean:  4,
				n:     2,
			},
			wantAns: []int{6, 6},
		},
		{
			name: "tc2",
			args: args{
				rolls: []int{1, 5, 6},
				mean:  3,
				n:     4,
			},
			wantAns: []int{3, 2, 2, 2},
		},
		{
			name: "tc3",
			args: args{
				rolls: []int{1, 2, 3, 4},
				mean:  6,
				n:     4,
			},
			wantAns: []int(nil),
		},
		{
			name: "tc4",
			args: args{
				rolls: []int{6, 3, 4, 3, 5, 3},
				mean:  1,
				n:     6,
			},
			wantAns: []int(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := missingRolls(tt.args.rolls, tt.args.mean, tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("missingRolls() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
