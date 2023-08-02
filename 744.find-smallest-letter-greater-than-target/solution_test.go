package findsmallestlettergreaterthantarget

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "tc1",
			args: args{
				letters: []byte{
					'c', 'f', 'j',
				},
				target: 'a',
			},
			want: 'c',
		},
		{
			name: "tc2",
			args: args{
				letters: []byte{
					'c', 'f', 'j',
				},
				target: 'c',
			},
			want: 'f',
		},
		{
			name: "tc3",
			args: args{
				letters: []byte{
					'x', 'x', 'y', 'y',
				},
				target: 'z',
			},
			want: 'x',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
