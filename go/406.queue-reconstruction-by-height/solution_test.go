package queuereconstructionbyheight

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "tc1",
			args: args{
				people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			},
			want: [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
		{
			name: "tc2",
			args: args{
				people: [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}},
			},
			want: [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
