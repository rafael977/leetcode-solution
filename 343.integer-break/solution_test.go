package integerbreak

import "testing"

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				n: 2,
			},
			wantAns: 1,
		},
		{
			name: "tc2",
			args: args{
				n: 10,
			},
			wantAns: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := integerBreak(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("integerBreak() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
