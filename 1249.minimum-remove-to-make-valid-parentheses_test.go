package main

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				s: "lee(t(c)o)de)",
			},
			want: "lee(t(c)o)de",
		},
		{
			name: "tc2",
			args: args{
				s: "a)b(c)d",
			},
			want: "ab(c)d",
		},
		{
			name: "tc3",
			args: args{
				s: "))((",
			},
			want: "",
		},
		{
			name: "tc4",
			args: args{
				s: "(lee(t(c)o)de",
			},
			want: "lee(t(c)o)de",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
