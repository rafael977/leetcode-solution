package parsingabooleanexpression

import "testing"

func Test_parseBoolExpr(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				expression: "&(|(f))",
			},
			want: false,
		},
		{
			name: "tc2",
			args: args{
				expression: "|(f,f,f,t)",
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				expression: "!(&(f,t))",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoolExpr(tt.args.expression); got != tt.want {
				t.Errorf("parseBoolExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
