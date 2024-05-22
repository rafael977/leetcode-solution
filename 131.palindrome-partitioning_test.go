package leetcodesolution

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]string
	}{
		{
			name: "tc1",
			args: args{
				s: "aab",
			},
			wantAns: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "tc2",
			args: args{
				s: "a",
			},
			wantAns: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := partition(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("partition() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
