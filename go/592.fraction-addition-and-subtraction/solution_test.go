package fractionadditionandsubtraction

import "testing"

func Test_fractionAddition(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "tc1",
			args: args{
				expression: "-1/2+1/2",
			},
			wantAns: "0/1",
		},
		{
			name: "tc2",
			args: args{
				expression: "-1/2+1/2+1/3",
			},
			wantAns: "1/3",
		},
		{
			name: "tc3",
			args: args{
				expression: "1/3-1/2",
			},
			wantAns: "-1/6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := fractionAddition(tt.args.expression); gotAns != tt.wantAns {
				t.Errorf("fractionAddition() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
