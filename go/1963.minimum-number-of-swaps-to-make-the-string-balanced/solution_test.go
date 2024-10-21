package minimumnumberofswapstomakethestringbalanced

import "testing"

func Test_minSwaps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				s: "][][",
			},
			want: 1,
		},
		{
			name: "tc2",
			args: args{
				s: "]]][[[",
			},
			want: 2,
		},
		{
			name: "tc3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "tc4",
			args: args{
				s: "[[][[[[][][[[[]]][][]]]]][[][[][][]][[[[]]][[]][[]][[]]]]]]]][]][]]][[]][[[[[[][[]][[[][]][[]]][",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.s); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
