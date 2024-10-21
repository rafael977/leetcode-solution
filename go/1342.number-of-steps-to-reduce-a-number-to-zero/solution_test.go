package numberofstepstoreduceanumbertozero

import (
	"testing"
)

func Test_numberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name      string
		args      args
		wantSteps int
	}{
		{
			name:      "tc1",
			args:      args{num: 14},
			wantSteps: 6,
		},
		{
			name:      "tc2",
			args:      args{num: 8},
			wantSteps: 4,
		},
		{
			name:      "tc3",
			args:      args{num: 123},
			wantSteps: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSteps := numberOfSteps(tt.args.num); gotSteps != tt.wantSteps {
				t.Errorf("numberOfSteps() = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}
