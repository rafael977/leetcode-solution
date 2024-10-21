package baseballgame

import "testing"

func Test_calPoints(t *testing.T) {
	type args struct {
		ops []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				ops: []string{"5", "2", "C", "D", "+"},
			},
			want: 30,
		},
		{
			name: "tc2",
			args: args{
				ops: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			},
			want: 27,
		},
		{
			name: "tc3",
			args: args{
				ops: []string{"1"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.args.ops); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
