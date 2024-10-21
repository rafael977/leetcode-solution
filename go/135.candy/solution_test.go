package candy

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				ratings: []int{1, 0, 2},
			},
			wantAns: 5,
		},
		{
			name: "tc2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := candy(tt.args.ratings); gotAns != tt.wantAns {
				t.Errorf("candy() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
