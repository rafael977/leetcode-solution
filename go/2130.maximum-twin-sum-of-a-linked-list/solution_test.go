package maximumtwinsumofalinkedlist

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_pairSum(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				head: BuildLinkedList("5,4,2,1"),
			},
			wantAns: 6,
		},
		{
			name: "tc2",
			args: args{
				head: BuildLinkedList("4,2,2,3"),
			},
			wantAns: 7,
		},
		{
			name: "tc3",
			args: args{
				head: BuildLinkedList("1,100000"),
			},
			wantAns: 100001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := pairSum(tt.args.head); gotAns != tt.wantAns {
				t.Errorf("pairSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
