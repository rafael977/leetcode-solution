package stringcompressionq

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			},
			want: 6,
		},
		{
			name: "tc2",
			args: args{
				chars: []byte{'a'},
			},
			want: 1,
		},
		{
			name: "tc3",
			args: args{
				chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
