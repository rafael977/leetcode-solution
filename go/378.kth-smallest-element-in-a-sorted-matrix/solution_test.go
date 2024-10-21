package kthsmallestelementinasortedmatrix

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				matrix: [][]int{
					{1, 5, 9},
					{10, 11, 13},
					{12, 13, 15},
				},
				k: 8,
			},
			want: 13,
		},
		{
			name: "tc2",
			args: args{
				matrix: [][]int{
					{-5},
				},
				k: 1,
			},
			want: -5,
		},
		{
			name: "tc3",
			args: args{
				matrix: [][]int{
					{1, 3, 5},
					{6, 7, 12},
					{11, 14, 14},
				},
				k: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
