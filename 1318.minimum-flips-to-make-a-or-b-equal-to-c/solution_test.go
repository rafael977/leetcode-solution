package minimumflipstomakeaorbequaltoc

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				a: 2,
				b: 6,
				c: 5,
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				a: 4,
				b: 2,
				c: 7,
			},
			wantAns: 1,
		},
		{
			name: "tc3",
			args: args{
				a: 1,
				b: 2,
				c: 3,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minFlips(tt.args.a, tt.args.b, tt.args.c); gotAns != tt.wantAns {
				t.Errorf("minFlips() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
