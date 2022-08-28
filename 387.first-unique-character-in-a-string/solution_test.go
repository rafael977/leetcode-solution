package firstuniquecharacterinastring

import "testing"

func Test_firstUniqChar(t *testing.T) {
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
			args:    args{s: "leetcode"},
			wantAns: 0,
		},
		{
			name:    "tc2",
			args:    args{s: "loveleetcode"},
			wantAns: 2,
		},
		{
			name:    "tc3",
			args:    args{s: "aabb"},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := firstUniqChar(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("firstUniqChar() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
