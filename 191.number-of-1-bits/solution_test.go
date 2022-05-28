package numberof1bits

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name  string
		args  args
		wantD int
	}{
		{
			name: "tc1",
			args: args{
				num: 11,
			},
			wantD: 3,
		},
		{
			name: "tc2",
			args: args{
				num: 128,
			},
			wantD: 1,
		},
		{
			name: "tc3",
			args: args{
				num: 4294967293,
			},
			wantD: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := hammingWeight(tt.args.num); gotD != tt.wantD {
				t.Errorf("hammingWeight() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}
