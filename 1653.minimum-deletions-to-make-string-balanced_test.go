package leetcodesolution

import "testing"

func Test_minimumDeletions(t *testing.T) {
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
				s: "aababbab",
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				s: "bbaaaaabb",
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumDeletions(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minimumDeletions() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
