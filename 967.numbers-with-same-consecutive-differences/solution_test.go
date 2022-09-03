package numberswithsameconsecutivedifferences

import (
	"reflect"
	"testing"
)

func Test_numsSameConsecDiff(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tc1",
			args: args{n: 3, k: 7},
			want: []int{181, 292, 707, 818, 929},
		},
		{
			name: "tc2",
			args: args{n: 2, k: 1},
			want: []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98},
		},
		{
			name: "tc3",
			args: args{n: 2, k: 0},
			want: []int{11, 22, 33, 44, 55, 66, 77, 88, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numsSameConsecDiff(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsSameConsecDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
