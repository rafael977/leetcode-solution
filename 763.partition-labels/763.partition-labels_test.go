package main

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				s: "ababcbacadefegdehijhklij",
			},
			wantAns: []int{9, 7, 8},
		},
		{
			name: "tc2",
			args: args{
				s: "eccbbbbdec",
			},
			wantAns: []int{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := partitionLabels(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("partitionLabels() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
