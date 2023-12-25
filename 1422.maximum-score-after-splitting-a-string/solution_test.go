package maximumscoreaftersplittingastring

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				s: "011101",
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				s: "00111",
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxScore(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
