package minimumbitflipstoconvertnumber

import "testing"

func Test_minBitFlips(t *testing.T) {
	type args struct {
		start int
		goal  int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				start: 10,
				goal:  7,
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				start: 3,
				goal:  4,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minBitFlips(tt.args.start, tt.args.goal); gotAns != tt.wantAns {
				t.Errorf("minBitFlips() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
