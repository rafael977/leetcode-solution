package optimalpartitionofstring

import "testing"

func Test_partitionString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				s: "abacaba",
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				s: "ssssss",
			},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := partitionString(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("partitionString() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
