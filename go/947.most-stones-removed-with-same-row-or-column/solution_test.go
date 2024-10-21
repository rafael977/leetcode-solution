package moststonesremovedwithsameroworcolumn

import "testing"

func Test_removeStones(t *testing.T) {
	type args struct {
		stones [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				stones: [][]int{
					{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2},
				},
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				stones: [][]int{
					{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2},
				},
			},
			wantAns: 3,
		},
		{
			name: "tc3",
			args: args{
				stones: [][]int{
					{0, 0},
				},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := removeStones(tt.args.stones); gotAns != tt.wantAns {
				t.Errorf("removeStones() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
