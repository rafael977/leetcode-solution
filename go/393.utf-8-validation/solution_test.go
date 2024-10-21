package utf8validation

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				data: []int{197, 130, 1},
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				data: []int{235, 140, 4},
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				data: []int{255},
			},
			want: false,
		},
		{
			name: "tc4",
			args: args{
				data: []int{145},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
