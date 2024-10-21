package leetcodesolution

import "testing"

func Test_countSeniors(t *testing.T) {
	type args struct {
		details []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				details: []string{
					"7868190130M7522", "5303914400F9211", "9273338290F4010",
				},
			},
			wantAns: 2,
		},
		{
			name: "tc2",
			args: args{
				details: []string{
					"1313579440F2036", "2921522980M5644",
				},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSeniors(tt.args.details); gotAns != tt.wantAns {
				t.Errorf("countSeniors() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
