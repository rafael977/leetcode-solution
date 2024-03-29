package satisfiabilityofequalityequations

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				equations: []string{"a==b", "b!=a"},
			},
			want: false,
		},
		{
			name: "tc2",
			args: args{
				equations: []string{"b==a", "a==b"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
