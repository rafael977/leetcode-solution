package minimummovestoequalarrayelementsii

import "testing"

func Test_minMoves2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				nums: []int{1, 10, 2, 9},
			},
			want: 16,
		},
		{
			name: "tc3",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.args.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
