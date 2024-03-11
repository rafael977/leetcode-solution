package leetcodesolution

import "testing"

func Test_customSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				order: "cba",
				s:     "abcd",
			},
			want: "cbad",
		},
		{
			name: "tc2",
			args: args{
				order: "bcafg",
				s:     "abcd",
			},
			want: "bcad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := customSortString(tt.args.order, tt.args.s); got != tt.want {
				t.Errorf("customSortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
