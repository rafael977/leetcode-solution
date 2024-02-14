package perfectsquares

import "testing"

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "tc1",
			args:    args{n: 12},
			wantAns: 3,
		},
		{
			name:    "tc2",
			args:    args{n: 13},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numSquares(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("numSquares() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
