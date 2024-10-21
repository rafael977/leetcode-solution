package uniquepaths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "tc1",
			args:    args{m: 3, n: 7},
			wantAns: 28,
		},
		{
			name:    "tc2",
			args:    args{m: 3, n: 2},
			wantAns: 3,
		},
		{
			name:    "tc3",
			args:    args{m: 13, n: 23},
			wantAns: 548354040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := uniquePaths(tt.args.m, tt.args.n); gotAns != tt.wantAns {
				t.Errorf("uniquePaths() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
