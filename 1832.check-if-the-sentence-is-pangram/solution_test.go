package checkifthesentenceispangram

import "testing"

func Test_checkIfPangram(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name    string
		args    args
		wantAns bool
	}{
		{
			name:    "tc1",
			args:    args{sentence: "thequickbrownfoxjumpsoverthelazydog"},
			wantAns: true,
		},
		{
			name:    "tc2",
			args:    args{sentence: "leetcode"},
			wantAns: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := checkIfPangram(tt.args.sentence); gotAns != tt.wantAns {
				t.Errorf("checkIfPangram() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
