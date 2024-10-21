package concatenationofconsecutivebinarynumbers

import "testing"

func Test_concatenatedBinary(t *testing.T) {
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
			args:    args{n: 1},
			wantAns: 1,
		},
		{
			name:    "tc2",
			args:    args{n: 3},
			wantAns: 27,
		},
		{
			name:    "tc3",
			args:    args{n: 12},
			wantAns: 505379714,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := concatenatedBinary(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("concatenatedBinary() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
