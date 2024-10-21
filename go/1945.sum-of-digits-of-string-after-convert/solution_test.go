package sumofdigitsofstringafterconvert

import "testing"

func Test_getLucky(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				s: "iiii",
				k: 1,
			},
			wantAns: 36,
		},
		{
			name: "tc2",
			args: args{
				s: "leetcode",
				k: 2,
			},
			wantAns: 6,
		},
		{
			name: "tc3",
			args: args{
				s: "zbax",
				k: 2,
			},
			wantAns: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getLucky(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("getLucky() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
