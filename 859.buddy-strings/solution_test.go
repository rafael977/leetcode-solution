package buddystrings

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				s:    "ab",
				goal: "ba",
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				s:    "ab",
				goal: "ab",
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				s:    "aa",
				goal: "aa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
