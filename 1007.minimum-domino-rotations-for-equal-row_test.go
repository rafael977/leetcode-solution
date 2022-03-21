package main

import "testing"

func Test_minDominoRotations(t *testing.T) {
	type args struct {
		tops    []int
		bottoms []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				tops:    []int{2, 1, 2, 4, 2, 2},
				bottoms: []int{5, 2, 6, 2, 3, 2},
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				tops:    []int{3, 5, 1, 2, 3},
				bottoms: []int{3, 6, 3, 3, 4},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDominoRotations(tt.args.tops, tt.args.bottoms); got != tt.want {
				t.Errorf("minDominoRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
