package decodedstringatindex

import "testing"

func Test_decodeAtIndex(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				s: "leet2code3",
				k: 10,
			},
			want: "o",
		},
		{
			name: "tc2",
			args: args{
				s: "ha22",
				k: 5,
			},
			want: "h",
		},
		{
			name: "tc3",
			args: args{
				s: "a2345678999999999999999",
				k: 1,
			},
			want: "a",
		},
		{
			name: "tc4",
			args: args{
				s: "y959q969u3hb22odq595",
				k: 222280369,
			},
			want: "y",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeAtIndex(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("decodeAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
