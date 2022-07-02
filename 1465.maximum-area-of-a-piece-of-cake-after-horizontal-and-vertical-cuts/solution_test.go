package maximumareaofapieceofcakeafterhorizontalandverticalcuts

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				h:              5,
				w:              4,
				horizontalCuts: []int{1, 2, 4},
				verticalCuts:   []int{1, 3},
			},
			wantAns: 4,
		},
		{
			name: "tc2",
			args: args{
				h:              5,
				w:              4,
				horizontalCuts: []int{3, 1},
				verticalCuts:   []int{1},
			},
			wantAns: 6,
		},
		{
			name: "tc3",
			args: args{
				h:              5,
				w:              4,
				horizontalCuts: []int{3},
				verticalCuts:   []int{3},
			},
			wantAns: 9,
		},
		{
			name: "tc4",
			args: args{
				h:              1000000000,
				w:              1000000000,
				horizontalCuts: []int{2},
				verticalCuts:   []int{2},
			},
			wantAns: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxArea(tt.args.h, tt.args.w, tt.args.horizontalCuts, tt.args.verticalCuts); gotAns != tt.wantAns {
				t.Errorf("maxArea() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
