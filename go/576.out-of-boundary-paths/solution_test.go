package outofboundarypaths

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				m:           2,
				n:           2,
				maxMove:     2,
				startRow:    0,
				startColumn: 0,
			},
			wantAns: 6,
		},
		{
			name: "tc2",
			args: args{
				m:           1,
				n:           3,
				maxMove:     3,
				startRow:    0,
				startColumn: 1,
			},
			wantAns: 12,
		},
		{
			name: "tc3",
			args: args{
				m:           10,
				n:           10,
				maxMove:     0,
				startRow:    5,
				startColumn: 5,
			},
			wantAns: 0,
		},
		{
			name: "tc4",
			args: args{
				m:           36,
				n:           5,
				maxMove:     50,
				startRow:    15,
				startColumn: 3,
			},
			wantAns: 390153306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); gotAns != tt.wantAns {
				t.Errorf("findPaths() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
