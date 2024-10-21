package separateblackandwhiteballs

import "testing"

func Test_minimumSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "tc1",
			args: args{s: "101"},
			want: 1,
		},
		{
			name: "tc2",
			args: args{s: "100"},
			want: 2,
		},
		{
			name: "tc3",
			args: args{s: "0111"},
			want: 0,
		},
		{
			name: "tc4",
			args: args{s: "111111111100100010"},
			want: 65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSteps(tt.args.s); got != tt.want {
				t.Errorf("minimumSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
