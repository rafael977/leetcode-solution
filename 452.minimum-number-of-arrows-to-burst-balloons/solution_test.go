package minimumnumberofarrowstoburstballoons

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			},
			wantAns: 4,
		},
		{
			name: "tc3",
			args: args{
				points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMinArrowShots(tt.args.points); gotAns != tt.wantAns {
				t.Errorf("findMinArrowShots() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
