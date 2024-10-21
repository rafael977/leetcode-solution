package minimumnumberoftapstoopentowateragarden

import "testing"

func Test_minTaps(t *testing.T) {
	type args struct {
		n      int
		ranges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				n:      5,
				ranges: []int{3, 4, 1, 1, 0, 0},
			},
			want: 1,
		},
		{
			name: "tc2",
			args: args{
				n:      3,
				ranges: []int{0, 0, 0, 0},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTaps(tt.args.n, tt.args.ranges); got != tt.want {
				t.Errorf("minTaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
