package pascalstriangleii

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				rowIndex: 3,
			},
			wantAns: []int{1, 3, 3, 1},
		},
		{
			name: "tc2",
			args: args{
				rowIndex: 0,
			},
			wantAns: []int{1},
		},
		{
			name: "tc3",
			args: args{
				rowIndex: 1,
			},
			wantAns: []int{1, 1},
		},
		{
			name: "tc4",
			args: args{
				rowIndex: 21,
			},
			wantAns: []int{1, 21, 210, 1330, 5985, 20349, 54264, 116280, 203490, 293930, 352716, 352716, 293930, 203490, 116280, 54264, 20349, 5985, 1330, 210, 21, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getRow(tt.args.rowIndex); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getRow() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
