package groupthepeoplegiventhegroupsizetheybelongto

import (
	"reflect"
	"testing"
)

func Test_groupThePeople(t *testing.T) {
	type args struct {
		groupSizes []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "tc1",
			args: args{
				groupSizes: []int{3, 3, 3, 3, 3, 1, 3},
			},
			wantAns: [][]int{
				{5}, {0, 1, 2}, {3, 4, 6},
			},
		},
		{
			name: "tc2",
			args: args{
				groupSizes: []int{2, 1, 3, 3, 3, 2},
			},
			wantAns: [][]int{
				{1}, {0, 5}, {2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := groupThePeople(tt.args.groupSizes); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("groupThePeople() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
