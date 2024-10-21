package detonatethemaximumbombs

import "testing"

func Test_maximumDetonation(t *testing.T) {
	type args struct {
		bombs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				bombs: [][]int{
					{2, 1, 3},
					{6, 1, 4},
				},
			},
			want: 2,
		},
		{
			name: "tc2",
			args: args{
				bombs: [][]int{
					{1, 1, 5},
					{10, 10, 5},
				},
			},
			want: 1,
		},
		{
			name: "tc3",
			args: args{
				bombs: [][]int{
					{1, 2, 3},
					{2, 3, 1},
					{3, 4, 2},
					{4, 5, 3},
					{5, 6, 4},
				},
			},
			want: 5,
		},
		{
			name: "tc4",
			args: args{
				bombs: [][]int{
					{855, 82, 158},
					{17, 719, 430},
					{90, 756, 164},
					{376, 17, 340},
					{691, 636, 152},
					{565, 776, 5},
					{464, 154, 271},
					{53, 361, 162},
					{278, 609, 82},
					{202, 927, 219},
					{542, 865, 377},
					{330, 402, 270},
					{720, 199, 10},
					{986, 697, 443},
					{471, 296, 69},
					{393, 81, 404},
					{127, 405, 17},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDetonation(tt.args.bombs); got != tt.want {
				t.Errorf("maximumDetonation() = %v, want %v", got, tt.want)
			}
		})
	}
}
