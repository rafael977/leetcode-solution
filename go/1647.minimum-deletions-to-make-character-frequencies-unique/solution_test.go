package minimumdeletionstomakecharacterfrequenciesunique

import "testing"

func Test_minDeletions(t *testing.T) {
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
				s: "aab",
			},
			wantAns: 0,
		},
		{
			name: "tc2",
			args: args{
				s: "aaabbbcc",
			},
			wantAns: 2,
		},
		{
			name: "tc3",
			args: args{
				s: "ceabaacb",
			},
			wantAns: 2,
		},
		{
			name: "tc4",
			args: args{
				s: "bbcebab",
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minDeletions(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("minDeletions() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
