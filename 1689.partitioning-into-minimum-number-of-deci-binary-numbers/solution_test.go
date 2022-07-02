package partitioningintominimumnumberofdecibinarynumbers

import "testing"

func Test_minPartitions(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				n: "32",
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				n: "82734",
			},
			wantAns: 8,
		},
		{
			name: "tc3",
			args: args{
				n: "27346209830709182346",
			},
			wantAns: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minPartitions(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("minPartitions() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
