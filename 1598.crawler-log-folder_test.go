package leetcodesolution

import "testing"

func Test_minOperations1598(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			},
			want: 3,
		},
		{
			name: "tc2",
			args: args{
				logs: []string{"d1/", "../", "../", "../"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations1598(tt.args.logs); got != tt.want {
				t.Errorf("minOperations1598() = %v, want %v", got, tt.want)
			}
		})
	}
}
