package countsortedvowelstrings

import "testing"

func Test_countVowelStrings(t *testing.T) {
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
			args: args{n: 1},
			want: 5,
		},
		{
			name: "tc2",
			args: args{n: 2},
			want: 15,
		},
		{
			name: "tc3",
			args: args{n: 33},
			want: 66045,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelStrings(tt.args.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
