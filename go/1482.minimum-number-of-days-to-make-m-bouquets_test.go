package leetcodesolution

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				bloomDay: []int{1, 10, 3, 10, 2},
				m:        3,
				k:        1,
			},
			want: 3,
		},
		{
			name: "tc2",
			args: args{
				bloomDay: []int{1, 10, 3, 10, 2},
				m:        3,
				k:        3,
			},
			want: -1,
		},
		{
			name: "tc3",
			args: args{
				bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
				m:        2,
				k:        3,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
