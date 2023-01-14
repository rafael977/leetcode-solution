package sametree

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				p: BuildTree("1,2,3"),
				q: BuildTree("1,2,3"),
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				p: BuildTree("1,2"),
				q: BuildTree("1,null,2"),
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				p: BuildTree("1,2,1"),
				q: BuildTree("1,1,2"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
