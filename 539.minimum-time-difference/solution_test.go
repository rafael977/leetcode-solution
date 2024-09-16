package minimumtimedifference

import "testing"

func Test_findMinDifference(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				timePoints: []string{
					"23:59", "00:00",
				},
			},
			wantAns: 1,
		},
		{
			name: "tc2",
			args: args{
				timePoints: []string{
					"00:00", "23:59", "00:00",
				},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMinDifference(tt.args.timePoints); gotAns != tt.wantAns {
				t.Errorf("findMinDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
