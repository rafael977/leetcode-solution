package jumpgameiv

import "testing"

func Test_minJumps(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				arr: []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				arr: []int{7},
			},
			wantAns: 0,
		},
		{
			name: "tc3",
			args: args{
				arr: []int{7, 6, 9, 6, 9, 6, 9, 7},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minJumps(tt.args.arr); gotAns != tt.wantAns {
				t.Errorf("minJumps() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
