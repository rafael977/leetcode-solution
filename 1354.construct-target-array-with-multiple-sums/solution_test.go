package constructtargetarraywithmultiplesums

import "testing"

func Test_isPossible(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				target: []int{9, 3, 5},
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				target: []int{1, 1, 1, 2},
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				target: []int{8, 5},
			},
			want: true,
		},
		{
			name: "tc4",
			args: args{
				target: []int{9, 9, 9},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.target); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
