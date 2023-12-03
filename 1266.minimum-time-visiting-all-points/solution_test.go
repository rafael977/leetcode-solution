package minimumtimevisitingallpoints

import "testing"

func Test_minTimeToVisitAllPoints(t *testing.T) {
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
				points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			},
			wantAns: 7,
		},
		{
			name: "tc2",
			args: args{
				points: [][]int{{3, 2}, {-2, 2}},
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minTimeToVisitAllPoints(tt.args.points); gotAns != tt.wantAns {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
