package longestsubstringwithoutrepeatingcharacters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
			args:    args{s: "abcabcbb"},
			wantAns: 3,
		},
		{
			name:    "tc2",
			args:    args{s: "bbbb"},
			wantAns: 1,
		},
		{
			name:    "tc3",
			args:    args{s: "pwwkew"},
			wantAns: 3,
		},
		{
			name:    "tc4",
			args:    args{s: "dvdf"},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := lengthOfLongestSubstring(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
