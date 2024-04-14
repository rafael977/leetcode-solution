package leetcodesolution

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			wantAns: 6,
		},
		{
			name: "tc2",
			args: args{
				matrix: [][]byte{
					{'0'},
				},
			},
			wantAns: 0,
		},
		{
			name: "tc3",
			args: args{
				matrix: [][]byte{
					{'1'},
				},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximalRectangle(tt.args.matrix); gotAns != tt.wantAns {
				t.Errorf("maximalRectangle() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
