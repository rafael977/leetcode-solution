package fruitintobaskets

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				fruits: []int{1, 2, 1},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				fruits: []int{0, 1, 2, 2},
			},
			wantAns: 3,
		},
		{
			name: "tc3",
			args: args{
				fruits: []int{1, 2, 3, 2, 2},
			},
			wantAns: 4,
		},
		{
			name: "tc4",
			args: args{
				fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			},
			wantAns: 5,
		},
		{
			name: "tc5",
			args: args{
				fruits: []int{1, 0, 1, 4, 1, 4, 1, 2, 3},
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := totalFruit(tt.args.fruits); gotAns != tt.wantAns {
				t.Errorf("totalFruit() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
