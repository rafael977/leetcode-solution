package bagoftokens

import "testing"

func Test_bagOfTokensScore(t *testing.T) {
	type args struct {
		tokens []int
		power  int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				tokens: []int{100},
				power:  50,
			},
			wantAns: 0,
		},
		{
			name: "tc2",
			args: args{
				tokens: []int{100, 200},
				power:  150,
			},
			wantAns: 1,
		},
		{
			name: "tc3",
			args: args{
				tokens: []int{100, 200, 300, 400},
				power:  200,
			},
			wantAns: 2,
		},
		{
			name: "tc4",
			args: args{
				tokens: []int{71, 55, 82},
				power:  54,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := bagOfTokensScore(tt.args.tokens, tt.args.power); gotAns != tt.wantAns {
				t.Errorf("bagOfTokensScore() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
