package minimumpenaltyforashop

import "testing"

func Test_bestClosingTime(t *testing.T) {
	type args struct {
		customers string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				customers: "YYNY",
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				customers: "NNNNN",
			},
			wantAns: 0,
		},
		{
			name: "tc3",
			args: args{
				customers: "YYYY",
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := bestClosingTime(tt.args.customers); gotAns != tt.wantAns {
				t.Errorf("bestClosingTime() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
