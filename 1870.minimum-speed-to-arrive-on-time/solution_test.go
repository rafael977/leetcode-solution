package minimumspeedtoarriveontime

import "testing"

func Test_minSpeedOnTime(t *testing.T) {
	type args struct {
		dist []int
		hour float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				dist: []int{1, 3, 2},
				hour: 6,
			},
			want: 1,
		},
		{
			name: "tc2",
			args: args{
				dist: []int{1, 3, 2},
				hour: 2.7,
			},
			want: 3,
		},
		{
			name: "tc3",
			args: args{
				dist: []int{1, 3, 2},
				hour: 1.9,
			},
			want: -1,
		},
		{
			name: "tc4",
			args: args{
				dist: []int{1, 1, 100000},
				hour: 2.01,
			},
			want: 10000000,
		},
		{
			name: "tc5",
			args: args{
				dist: []int{5, 3, 4, 6, 2, 2, 7},
				hour: 10.92,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSpeedOnTime(tt.args.dist, tt.args.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
