package calculatemoneyinleetcodebank

import "testing"

func Test_totalMoney(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "tc1",
			args:    args{n: 4},
			wantAns: 10,
		},
		{
			name:    "tc2",
			args:    args{n: 10},
			wantAns: 37,
		},
		{
			name:    "tc3",
			args:    args{n: 20},
			wantAns: 96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := totalMoney(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("totalMoney() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
