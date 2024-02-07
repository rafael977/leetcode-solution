package amountoftimeforbinarytreetobeinfected

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_amountOfTime(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				root:  BuildTree("1,5,3,null,4,10,6,9,2"),
				start: 4,
			},
			want: 4,
		},
		{
			name: "tc2",
			args: args{
				root:  BuildTree("1"),
				start: 1,
			},
			want: 0,
		},
		{
			name: "tc3",
			args: args{
				root:  BuildTree("1,null,2,3,4,null,5"),
				start: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := amountOfTime(tt.args.root, tt.args.start); got != tt.want {
				t.Errorf("amountOfTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
