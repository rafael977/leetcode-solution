package shortencodingofwords

import "testing"

func Test_minimumLengthEncoding(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				words: []string{"me", "time", "me", "bell"},
			},
			wantAns: 10,
		},
		{
			name: "tc2",
			args: args{
				words: []string{"t"},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumLengthEncoding(tt.args.words); gotAns != tt.wantAns {
				t.Errorf("minimumLengthEncoding() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
