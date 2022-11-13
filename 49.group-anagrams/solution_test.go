package groupanagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{
			name: "tc1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			wantAns: [][]string{
				{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := groupAnagrams(tt.args.strs); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("groupAnagrams() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
