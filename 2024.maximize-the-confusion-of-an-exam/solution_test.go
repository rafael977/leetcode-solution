package maximizetheconfusionofanexam

import "testing"

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				answerKey: "TTFF",
				k:         2,
			},
			want: 4,
		},
		{
			name: "tc2",
			args: args{
				answerKey: "TFFT",
				k:         1,
			},
			want: 3,
		},
		{
			name: "tc3",
			args: args{
				answerKey: "TTFTTFTT",
				k:         1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
