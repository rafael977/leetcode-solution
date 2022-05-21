package criticalconnectionsinanetwork

import (
	"reflect"
	"testing"
)

func Test_criticalConnections(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "tc1",
			args: args{
				n:           4,
				connections: [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}},
			},
			want: [][]int{{1, 3}},
		},
		{
			name: "tc2",
			args: args{
				n:           2,
				connections: [][]int{{0, 1}},
			},
			want: [][]int{{0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := criticalConnections(tt.args.n, tt.args.connections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("criticalConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}
