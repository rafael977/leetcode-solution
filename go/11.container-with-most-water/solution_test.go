package containerwithmostwater

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{
			name: "tc1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			wantMax: 49,
		},
		{
			name: "tc2",
			args: args{
				height: []int{1, 1},
			},
			wantMax: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxArea(tt.args.height); gotMax != tt.wantMax {
				t.Errorf("maxArea() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
