package leetcodesolution

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "tc1",
			args: args{
				num: 123,
			},
			wantAns: "One Hundred Twenty Three",
		},
		{
			name: "tc2",
			args: args{
				num: 12345,
			},
			wantAns: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			name: "tc3",
			args: args{
				num: 1234567,
			},
			wantAns: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			name: "tc5",
			args: args{
				num: 1000000,
			},
			wantAns: "One Million",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberToWords(tt.args.num); gotAns != tt.wantAns {
				t.Errorf("numberToWords() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
