package minimumscoreofapathbetweentwocities

import "testing"

func Test_minScore(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				n:     4,
				roads: [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}},
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				n:     4,
				roads: [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minScore(tt.args.n, tt.args.roads); gotAns != tt.wantAns {
				t.Errorf("minScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
