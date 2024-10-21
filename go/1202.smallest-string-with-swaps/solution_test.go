package smalleststringwithswaps

import "testing"

func Test_smallestStringWithSwaps(t *testing.T) {
	type args struct {
		s     string
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}},
			},
			want: "bacd",
		},
		{
			name: "tc2",
			args: args{
				s:     "dcab",
				pairs: [][]int{{0, 3}, {1, 2}, {0, 2}},
			},
			want: "abcd",
		},
		{
			name: "tc3",
			args: args{
				s:     "cba",
				pairs: [][]int{{0, 1}, {1, 2}},
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStringWithSwaps(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("smallestStringWithSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
