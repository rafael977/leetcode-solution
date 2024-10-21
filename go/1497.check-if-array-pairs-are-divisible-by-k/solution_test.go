package checkifarraypairsaredivisiblebyk

import "testing"

func Test_canArrange(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9},
				k:   5,
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6},
				k:   7,
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6},
				k:   10,
			},
			want: false,
		},
		{
			name: "tc4",
			args: args{
				arr: []int{-1, 1, -2, 2, -3, 3, -4, 4},
				k:   3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canArrange(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("canArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
