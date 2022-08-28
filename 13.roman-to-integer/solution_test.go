package romantointeger

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "tc1",
			args:    args{s: "III"},
			wantAns: 3,
		},
		{
			name:    "tc2",
			args:    args{s: "LVIII"},
			wantAns: 58,
		},
		{
			name:    "tc3",
			args:    args{s: "MCMXCIV"},
			wantAns: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := romanToInt(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("romanToInt() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
