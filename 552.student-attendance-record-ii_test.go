package leetcodesolution

import "testing"

func Test_checkRecord(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{n: 2},
			want: 8,
		},
		{
			name: "tc2",
			args: args{n: 1},
			want: 3,
		},
		{
			name: "tc3",
			args: args{n: 10101},
			want: 183236316,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord(tt.args.n); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
