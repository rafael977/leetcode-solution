package dividetwointegers

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name         string
		args         args
		wantQuotient int
	}{
		{
			name: "tc1",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			wantQuotient: 3,
		},
		{
			name: "tc2",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			wantQuotient: -2,
		},
		{
			name: "tc3",
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			wantQuotient: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotQuotient := divide(tt.args.dividend, tt.args.divisor); gotQuotient != tt.wantQuotient {
				t.Errorf("divide() = %v, want %v", gotQuotient, tt.wantQuotient)
			}
		})
	}
}
